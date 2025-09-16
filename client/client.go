package main

import (
	"net/http"
)

func main() {
	client := NewClient()
}

func NewClient() *http.Client {
	tr := &http.Transport{
		MaxConnsPerHost:   1,
		DisableKeepAlives: true,
	}

	return &http.Client{
		Transport: tr,
	}
}
}
