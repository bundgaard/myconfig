package myconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
Configuration handles following types from JSON
To read a configuration it needs to be called by your application name . json

Recipients   []string `"json:Recipients"`
	Body         string   `"json:Body"`
	Subject      string   `"json:Subject"`
	Filename     string   `"json:Filename"`
	Weekly       bool     `"json:Weekly"`
	Monthly      bool     `"json:Monthly"`
	Daily        bool     `"json:Daily"`
	SQLFile      string   `"json:SQLFile"`
	MaiHost      string   `"json:MailHost"`
	MailPort     int      `"json:MailPort"`
	MailUsername string   `"json:MailUsername"`
	MailPassword string   `"json:MailPassword"`
*/

type Configuration struct {
	Recipients   []string `"json:Recipients"`
	Body         string   `"json:Body"`
	Subject      string   `"json:Subject"`
	Filename     string   `"json:Filename"`
	Weekly       bool     `"json:Weekly"`
	Monthly      bool     `"json:Monthly"`
	Daily        bool     `"json:Daily"`
	SQLFile      string   `"json:SQLFile"`
	MaiHost      string   `"json:MailHost"`
	MailPort     int      `"json:MailPort"`
	MailUsername string   `"json:MailUsername"`
	MailPassword string   `"json:MailPassword"`
}

func LoadConfiguration() (Configuration, error) {
	var (
		appName string
		config  Configuration
	)

	appName = fmt.Sprintf("%s.json", os.Args[0])

	data, err := ioutil.ReadFile(appName)
	if err != nil {
		return Configuration{}, err
	}
	if err = json.Unmarshal(data, &config); err != nil {
		return Configuration{}, err
	}

	return config, nil
}
