package service

import (
	"fmt"
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

	// repositoryName, repositoryOwner := service.InputData()
	repositoryName, repositoryOwner := "Main-WebApp", "SIC-Unud"

	responseData, err := serviceObject.SendRequest(&service.SendRequestParams{
		GithubToken:     params.GithubToken,
		RepositoryName:  repositoryName,
		RepositoryOwner: repositoryOwner,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	for name, v := range *responseData {
		fmt.Println(name, v)
	}
}
