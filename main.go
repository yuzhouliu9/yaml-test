package main

import (
	"io/ioutil"

	"github.com/buildkite/yaml"
	//"gopkg.in/yaml.v3"
)

type Config struct {
	Schema string `yaml:"_schema"`
	Bar    []string
}

func main() {

	cfg := Config{}
	source, err := ioutil.ReadFile("in.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(source), &cfg)
	if err != nil {
		panic(err)
	}
	y, _ := yaml.Marshal(&cfg)
	err = ioutil.WriteFile("out.yaml", y, 0644)
}
