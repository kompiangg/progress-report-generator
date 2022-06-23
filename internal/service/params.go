package service

import "github.com/kompiangg/report-generator/internal/dto"

type RepositoryMetadata struct {
	RepositoryName  string
	RepositoryOwner string
}

type SendRequestParams struct {
	GithubToken string
	RepositoryMetadata
}

type GenerateFileParams struct {
	ReportData *dto.RepositoryData
	RepositoryMetadata
}
