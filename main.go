package main

import (
	"fmt"

	"github.com/kompiangg/report-generator/cmd/service"
	"github.com/kompiangg/report-generator/pkg/config"
	httpClientPkg "github.com/kompiangg/report-generator/pkg/http/client"
)

func main() {
	fmt.Printf("Program Started\n\n")

	config.Init("./config/.env")
	confValue := config.GetConfig()

	httpClientPkg.NewHttpClient()
	httpClient := httpClientPkg.GetHttpClient()

	paramsInitService := &service.ParamsInitService{
		GithubToken: confValue.GithubToken,
		HttpClient:  httpClient,
	}

	service.Start(paramsInitService)
}
