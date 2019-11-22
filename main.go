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
	Discover:    false,
	Devices: []sdk.Device{
		sdk.Device{
			Name:           "get",
			DefaultTrigger: "",
			DefaultAction:  "get",
			Config: []sdk.DeviceConfig{
				sdk.DeviceConfig{
					Name: "Link",
					Type: "string",
				},
			},
			Triggers: []sdk.Trigger{},
			Actions:  []string{"get"},
		},
		sdk.Device{
			Name:           "post",
			DefaultTrigger: "",
			DefaultAction:  "post",
			Config: []sdk.DeviceConfig{
				sdk.DeviceConfig{
					Name: "Link",
					Type: "string",
				},
			},
			Triggers: []sdk.Trigger{},
			Actions:  []string{"post"},
		},
	},
	Actions: []sdk.Action{
		sdk.Action{
			Name:   "get",
			Fields: []sdk.Field{},
		},
		sdk.Action{
			Name: "post",
			Fields: []sdk.Field{
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

// ConfigDevice define config for device
type ConfigDevice struct {
	Link string
}

// Params define action parameters available
type Params struct {
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
	if string(params) == "" {
		fmt.Println("Params must be provided")
	}
	if string(config) == "" {
		fmt.Println("Configs must be provided")
	}

	var conf ConfigDevice
	var marshalParam Params

	json.Unmarshal(config, &conf)
	json.Unmarshal(params, &marshalParam)

	// use name to call actions
	switch name {
	case "get":
		Get(conf.Link)
	case "post":
		Post(conf.Link, marshalParam.CtnType, marshalParam.Values)
	default:
		return
	}
}

// OnStop close connection
func OnStop() {
}
