//go:build !wasm

package config

import (
	"log"
	"os"
)

func readFile() ([]byte, error) {
	log.Println("reading config file")

	return os.ReadFile("config/config.yml")
}
