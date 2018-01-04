package main

import (
	"gopkg.in/yaml.v2"
)

type SayConfig struct {
	Voice struct {
		Ja string `yaml:"ja"`
		En string `yaml:"en"`
	} `yaml:"voice"`
	Volume float64 `yaml:"volume"`
	Rate   int     `yaml:"rate"`
}

func parseConfig(data []byte) (SayConfig, error) {
	var config SayConfig

	err := yaml.Unmarshal(data, &config)
	return config, err
}
