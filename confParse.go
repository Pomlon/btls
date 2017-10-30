package main

//TaskDefinition thanks https://mholt.github.io/json-to-go/ !!
type TaskDefinitions []struct {
	TaskName    string   `json:"TaskName"`
	TaskDescr   string   `json:"TaskDescr"`
	BeforeBuild []string `json:"BeforeBuild"`
	AfterBuild  []string `json:"AfterBuild"`
	WatchPaths  []string `json:"WatchPaths"`
	RunCommand  string   `json:"RunCommand"`
}
