package config

import (
	"../logger"
	"github.com/BurntSushi/toml"
)

var C map[string]interface{}

var defaultConfig string = `
lang = "en_US"
`

func init() {
	C = make(map[string]interface{})
	// Decode built in default config
	_, err := toml.Decode(defaultConfig, C)
	if err != nil {
		logger.Err("Error decoding default config blob")
	}

	// TODO: Load config files from /etc/yokel/config and ~/.config/yokel/config here

	// Load local dir config
	_, err = toml.DecodeFile("config", C)
	if err != nil {
		logger.Info("No config found at `./config'.")
	}
}
