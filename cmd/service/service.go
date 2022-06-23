package service

import (
	"log"
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

	repositoryMetadata := service.RepositoryMetadata{
		RepositoryName:  repositoryName,
		RepositoryOwner: repositoryOwner,
	}

	responseData, err := serviceObject.SendRequest(&service.SendRequestParams{
		GithubToken:        params.GithubToken,
		RepositoryMetadata: repositoryMetadata,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	serviceObject.GenerateFile(&service.GenerateFileParams{
		ReportData:         responseData,
		RepositoryMetadata: repositoryMetadata,
	})
}
