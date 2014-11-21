// Package
package config

import (
	"github.com/BurntSushi/toml"
	"github.com/SamWhited/logger"
	"github.com/mitchellh/go-homedir"
)

var C map[string]interface{}

// Initializes the config map if it hasn't been already.
func init() {
	if C == nil {
		C = make(map[string]interface{})
	}
}

// Loads a TOML blob into the config map provided by the package.
func LoadBlob(blob string) error {
	_, err := toml.Decode(blob, C)
	return err
}

// Loads a TOML file into the config map provided by the package.
func LoadFile(path string) error {
	_, err := toml.DecodeFile(path, C)
	return err
}

// Loads an entire project's config (system level, user level, and local config) into the map provided by the package.
// Logs any errors that occur using the `github.com/SamWhited/logger` package.
// TODO: Don't just log errors now that this is a separate package; push them to a channel or something and let the user
// handle them.
func LoadProjectConfig(project string) {
	// Load system-level config
	err := LoadFile("/etc/" + ProjectName + "/config")
	if err != nil {
		logger.Info("No config found at `/etc/" + ProjectName + "/config'.")
	}

	// Load user-level config
	home, err := homedir.Dir()
	if err != nil {
		logger.Info("No home directory for current user; not loading user-level config")
	} else {
		err = LoadFile(home + "/.config/" + ProjectName + "/config")
		if err != nil {
			logger.Info("No config found at `" + home + "/.config/" + ProjectName + "/config'.")
		}
	}

	// Load local config
	err = LoadFile("config")
	if err != nil {
		logger.Info("No config found at `./config'.")
	}
}
