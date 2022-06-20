package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kompiangg/report-generator/internal/constant"
	"github.com/kompiangg/report-generator/internal/dto"
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

func (s *service) SendRequest(githubToken string, repositoryName string, repositoryOwner string) *dto.RepositoryData {
	GET_ISSUES["query"] = fmt.Sprintf(GET_ISSUES["query"], repositoryName, repositoryOwner)

	jsonValue, err := json.Marshal(GET_ISSUES)
	if err != nil {
		panic("ERROR: error while encode the JSON")
	}

	request, err := http.NewRequest(http.MethodPost, constant.GITHUB_GRAPHQL_ENDPOINT, bytes.NewBuffer(jsonValue))
	if err != nil {
		panic("ERROR: error while creating new request")
	}

	request.Header.Set("Authorization", "Bearer "+githubToken)
	request.Header.Set("Content-Type", "application/json")

	response, err := s.httpClient.Do(request)
	if err != nil {
		panic("ERROR: error while hitting the GraphQL Github API")
	}

	defer response.Body.Close()

	bodyReader, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("ERROR: error while reading response body")
	}

	var bodyResponseValue interface{}

	err = json.Unmarshal(bodyReader, &bodyResponseValue)
	if err != nil {
		panic("ERROR: error while unmarshal response body")
	}

	return dto.NewRepositoryData(bodyResponseValue)
}
