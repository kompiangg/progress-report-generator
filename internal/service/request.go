package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kompiangg/report-generator/internal/constant"
	"github.com/kompiangg/report-generator/internal/dto"
	"github.com/kompiangg/report-generator/pkg/errors"
)

func (s *service) SendRequest(params *SendRequestParams) (*dto.RepositoryData, error) {
	GET_ISSUES["query"] = fmt.Sprintf(GET_ISSUES["query"], params.RepositoryName, params.RepositoryOwner)

	jsonValue, err := json.Marshal(GET_ISSUES)
	if err != nil {
		return nil, errors.ErrMarshalJSON
	}

	request, err := http.NewRequest(http.MethodPost, constant.GITHUB_GRAPHQL_ENDPOINT, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, errors.ErrCreatingNewRequest
	}

	request.Header.Set("Authorization", "Bearer "+params.GithubToken)
	request.Header.Set("Content-Type", "application/json")

	responseChan := make(chan *http.Response)

	fmt.Println("Sending request start")
	fmt.Println("Waiting for response")
	fmt.Printf("Progress ")

	go func() {
		response, err := s.httpClient.Do(request)
		if err != nil {
			responseChan <- nil
		}
		responseChan <- response
	}()

	var response *http.Response

	for resCheck := false; !resCheck; {
		select {
		case response = <-responseChan:
			if response == nil {
				return nil, errors.ErrHittingGraphQL
			}
			resCheck = true
			fmt.Printf("\nResponse received\n\n")
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Printf(". ")
		}
	}

	err = isOK(response.StatusCode)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyReader, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.ErrReadingResponseBody
	}

	var bodyResponseValue interface{}

	err = json.Unmarshal(bodyReader, &bodyResponseValue)
	if err != nil {
		return nil, errors.ErrUnmarshalJSON
	}

	return dto.NewRepositoryData(bodyResponseValue), nil
}

func isOK(statusCode int) error {
	if statusCode == 200 {
		return nil
	} else if statusCode == 401 {
		fmt.Println("ERROR: We get unauthorized error")
		fmt.Println("ERROR: Please restart the program and gimme a valid Authorization Token")
		return errors.ErrUnauthorizedRequest
	}
	return nil
}
