package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kompiangg/report-generator/internal/constant"
	"github.com/kompiangg/report-generator/internal/dto"
	"github.com/kompiangg/report-generator/pkg/errors"
)

var (
	GET_ISSUES = map[string]string{
		"query": `{
			repository(name: "%s", owner: "%s") {
					issues (states: CLOSED, first: 100) {
							nodes {
									title
									assignees (first: 100) {
											nodes {
													name
											}
									}
							}
					}
			}
		}`}
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

	response, err := s.httpClient.Do(request)
	if err != nil {
		return nil, errors.ErrHittingGraphQL
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
