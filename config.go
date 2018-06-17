package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Files []*File `yaml:"files"`
}

type File struct {
	Path         string         `yaml:"path"`
	Replacements []*Replacement `yaml:"replacements"`
}

type Replacement struct {
	Regex       string `yaml:"regex"`
	Replacement string `yaml:"replacement"`
	Group       uint8  `yaml:"group,omitempty"`
}

func loadConfigFromFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
