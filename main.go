package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {}

// Params define actions parameters available
type Params struct {
	Link    string
	CtnType string
	Values  string
}

var client http.Client

// OnStart create http client
func OnStart() {
	client = http.Client{
		Timeout: time.Second * 5,
	}
}

// CallAction call functions from actions
func CallAction(name string, params []byte) {
	if string(params) == "" {
		fmt.Println("Params must be provided")
		return
	}

	// declare parameters
	var req Params

	// unmarshal parameters to use in actions
	err := json.Unmarshal(params, &req)
	if err != nil {
		fmt.Println(err)
	}

	// use name to call actions
	switch name {
	case "get":
		Get(req.Link)
	case "post":
		Post(req.Link, req.CtnType, req.Values)
	default:
		return
	}
}
