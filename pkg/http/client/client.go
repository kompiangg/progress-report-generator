package client

import "net/http"

var client *http.Client

func GetHttpClient() *http.Client {
	return client
}

func NewHttpClient() *http.Client {
	client = &http.Client{}
	return GetHttpClient()
}
