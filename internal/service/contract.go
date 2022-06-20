package service

import (
	"net/http"

	"github.com/kompiangg/report-generator/internal/dto"
)

type ServiceContract interface {
	SendRequest(params *SendRequestParams) (*dto.RepositoryData, error)
	InputData() (repositoryName string, repositoryOwner string)
	GenerateFile(params *GenerateFileParams) error
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
