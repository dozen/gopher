package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfData struct {
	Password           string `yaml:"password"`
	Server             string `yaml:"server"`
	Channel            string `yaml:"channel"`
	GitWorkDir         string `yaml:"git_work_dir"`
	PullRequestComment string `yaml:"pull_request_comment"`
	MirageUrl          string `yaml:"mirage_url"`
	DockerImage        string `yaml:"docker_image"`
}

// conf.ymlをロードし構造体へ格納
func LoadConfig() *ConfData {
	buf, err := ioutil.ReadFile("config/config_local.yml")

	d := ConfData{}
	if err != nil {
		buf, err := ioutil.ReadFile("config/config.yml")
		if err != nil {
			panic(err)
		}

		if err := yaml.Unmarshal(buf, &d); err != nil {
			panic(err)
		}
		return &d
	}

	if err := yaml.Unmarshal(buf, &d); err != nil {
		panic(err)
	}

	return &d
}