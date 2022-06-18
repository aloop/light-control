package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Request struct {
}

func (r *Request) MakeRequest(method string, url string, data []byte) ([]byte, int) {
	client := &http.Client{
		Timeout: time.Duration(1) * time.Second,
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+config.API.AuthToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body, resp.StatusCode
}

func (r *Request) Get(url string) ([]byte, int) {
	return r.MakeRequest("GET", url, nil)
}

func (r *Request) Post(url string, data []byte) ([]byte, int) {
	return r.MakeRequest("POST", url, data)
}
