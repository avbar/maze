//go:build wasm

package config

import (
	"errors"
	"log"

	"github.com/avbar/maze/config"
)

func readFile() ([]byte, error) {
	log.Println("reading config for wasm")

	rawYAML := config.ConfigFile
	if rawYAML == nil {
		return nil, errors.New("empty config")
	}

	return rawYAML, nil
}
