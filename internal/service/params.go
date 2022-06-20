package service

import "github.com/kompiangg/report-generator/internal/dto"

type SendRequestParams struct {
	GithubToken     string
	RepositoryName  string
	RepositoryOwner string
}

type GenerateFileParams struct {
	ReportData *dto.RepositoryData
}
