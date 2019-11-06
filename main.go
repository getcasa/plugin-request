package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/getcasa/sdk"
)

func main() {}

// Config set the plugin config
var Config = sdk.Configuration{
	Name:        "request",
	Version:     "1.0.0",
	Author:      "amoinier",
	Description: "request",
	Main:        "request",
	FuncData:    "onData",
	Discover:    true,
	Triggers:    []sdk.Trigger{},
	Actions: []sdk.Action{
		sdk.Action{
			Name: "get",
			Fields: []sdk.Field{
				sdk.Field{
					Name:   "Link",
					Type:   "string",
					Config: true,
				},
			},
		},
		sdk.Action{
			Name: "post",
			Fields: []sdk.Field{
				sdk.Field{
					Name:   "Link",
					Type:   "string",
					Config: true,
				},
				sdk.Field{
					Name:   "CtnType",
					Type:   "string",
					Config: true,
				},
				sdk.Field{
					Name:   "Values",
					Type:   "string",
					Config: true,
				},
			},
		},
	},
}

// Params define actions parameters available
type Params struct {
	Link    string
	CtnType string
	Values  string
}

var client http.Client

// Init
func Init() []byte {
	return []byte{}
}

// OnStart create http client
func OnStart(config []byte) {
	client = http.Client{
		Timeout: time.Second * 5,
	}
}

// CallAction call functions from actions
func CallAction(name string, params []byte, config []byte) {
	if string(config) == "" {
		fmt.Println("Params must be provided")
		return
	}

	// declare parameters
	var req Params

	// unmarshal parameters to use in actions
	err := json.Unmarshal(config, &req)
	if err != nil {
		fmt.Println(err)
		return
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

// OnStop close connection
func OnStop() {
}
