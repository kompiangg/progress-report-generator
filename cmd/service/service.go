package service

import (
	"fmt"
	"net/http"

	"github.com/kompiangg/report-generator/internal/service"
)

type ParamsInitService struct {
	GithubToken string
	HttpClient  *http.Client
}

func Start(params *ParamsInitService) {
	serviceParams := &service.ServiceParams{
		HttpClient: params.HttpClient,
	}
	serviceObject := service.NewService(serviceParams)

	repositoryName, repositoryOwner := service.InputData()

	responseData := serviceObject.SendRequest(&service.SendRequestParams{
		GithubToken:     params.GithubToken,
		RepositoryName:  repositoryName,
		RepositoryOwner: repositoryOwner,
	})

	for name, v := range *responseData {
		fmt.Println(name, v)
	}
}
