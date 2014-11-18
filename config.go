package config

import (
	"../logger"
	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
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

	// Load system-level config
	_, err = toml.DecodeFile("/etc/yokel/config", C)
	if err != nil {
		logger.Info("No config found at `/etc/yokel/config'.")
	}

	// Load user-level config
	home, err := homedir.Dir()
	if err != nil {
		logger.Info("No home directory for current user; not loading user-level config")
	} else {
		_, err = toml.DecodeFile(home+"/.config/yokel/config", C)
		if err != nil {
			logger.Info("No config found at `" + home + "/.config/yokel/config'.")
		}
	}

	// Load local config
	_, err = toml.DecodeFile("config", C)
	if err != nil {
		logger.Info("No config found at `./config'.")
	}
}
