//go:build wasm

package config

import (
	_ "embed"
)

//go:embed config.yml
var ConfigFile []byte
