package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var url = "http://localhost:8080"

func main() {
	client := NewClient()
	ConnectClientToUrl(client)
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

func ConnectClientToUrl(cl *http.Client) {
	resp, err := cl.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Default().Println(err)
	}
	if !checkResponse(string(body)) {
		fmt.Println("Wrong response!")
		return
	}
	fmt.Println("Successful one-time connection")
}

func checkResponse(resp string) bool {
	return resp == "OK\n"
}
