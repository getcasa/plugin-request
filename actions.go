package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get is a function to do a get request
func Get(link string) ([]byte, error) {
	fmt.Println("Do get request")
	req, err := http.NewRequest("GET", link, nil)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

// Post is a function to do a post request
func Post(link string, ctnType string, values string) ([]byte, error) {
	if ctnType == "" {
		ctnType = "application/json"
	}
	fmt.Println("Do post request")
	req, err := http.NewRequest("POST", link, bytes.NewBuffer([]byte(values)))
	req.Header.Set("Content-Type", ctnType)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	return body, err
}
