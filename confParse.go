package main

import (
	"encoding/json"
	"io/ioutil"
)

//TaskDefinitions thanks https://mholt.github.io/json-to-go/ !!
type TaskDefinitions []struct {
	TaskName    string   `json:"TaskName"`
	TaskDescr   string   `json:"TaskDescr"`
	Build       []string `json:"Build"`
	AfterBuild  []string `json:"AfterBuild"`
	WatchPath   string   `json:"WatchPath"`
	IgnorePaths []string `json:"IgnorePaths"`
	RunCommand  string   `json:"RunCommand"`
}

//ParseConfig parses task configs from buildConfig.json
func ParseConfig() TaskDefinitions {

	config, err := ioutil.ReadFile("buildConfig.json")
	if err != nil {
		panic("No config no fun mate!")
	}

	taskDefs := TaskDefinitions{}
	json.Unmarshal(config, &taskDefs)

	return taskDefs
}
