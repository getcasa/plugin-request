package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Link is a valid URL

var client http.Client

func main() {
	var cl = http.Client{
		Timeout: time.Second * 5,
	}

	client = cl
}

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
	fmt.Println("Do post request")
	data := url.Values{}
	req, err := http.NewRequest("POST", link, strings.NewReader(data.Encode()))
	// req.Header.Add("Authorization", "")
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
