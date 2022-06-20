package service

import (
	"net/http"

	"github.com/kompiangg/report-generator/internal/dto"
)

type ServiceContract interface {
	SendRequest(githubToken string, repositoryName string, repositoryOwner string) *dto.RepositoryData
	InputData() (repositoryName string, repositoryOwner string)
}

type service struct {
	httpClient *http.Client
}

type ServiceParams struct {
	HttpClient *http.Client
}

func NewService(params *ServiceParams) *service {
	return &service{
		httpClient: params.HttpClient,
	}
}
