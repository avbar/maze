package config

import (
	"os"

	"github.com/avbar/maze/internal/settings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ScreenWidth  int               `yaml:"screen_width"`
	ScreenHeight int               `yaml:"screen_height"`
	Settings     settings.Settings `yaml:"settings"`
}

func MustLoad() *Config {
	rawYAML, err := os.ReadFile("config/config.yml")
	if err != nil {
		panic("cannot read config file: " + err.Error())
	}

	var cfg Config
	err = yaml.Unmarshal(rawYAML, &cfg)
	if err != nil {
		panic("cannot parse config file: " + err.Error())
	}

	return &cfg
}
