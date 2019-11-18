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
	Description: "Send request",
	Devices: []sdk.Device{
		sdk.Device{
			Name:           "get",
			DefaultTrigger: "",
			DefaultAction:  "",
			Triggers:       []sdk.Trigger{},
			Actions:        []string{"get"},
		},
	},
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

// ConfigDevice define actions parameters available
type ConfigDevice struct {
	Link    string
	CtnType string
	Values  string
}

var client http.Client

// OnStart create http client
func OnStart(config []byte) {
	client = http.Client{
		Timeout: time.Second * 5,
	}
}

// CallAction call functions from actions
func CallAction(physicalID string, name string, params []byte, config []byte) {
	if string(config) == "" {
		fmt.Println("Params must be provided")
		return
	}

	// declare parameters
	var conf ConfigDevice

	// unmarshal parameters to use in actions
	err := json.Unmarshal(config, &conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// use name to call actions
	switch name {
	case "get":
		Get(conf.Link)
	case "post":
		Post(conf.Link, conf.CtnType, conf.Values)
	default:
		return
	}
}

// OnStop close connection
func OnStop() {
}
