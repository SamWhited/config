// Package
package config

import (
	"github.com/BurntSushi/toml"
	"github.com/SamWhited/logger"
	"github.com/mitchellh/go-homedir"
)

var (
	C             map[string]interface{}
	DefaultConfig string
	ProjectName   string
)

// Initializes the config map if it hasn't already been initialized before package load.
func init() {
	if C == nil {
		C = make(map[string]interface{})
	}
	if DefaultConfig != nil && DefaultConfig != "" {
		LoadBlob(DefaultConfig)
	}
	if ProjectName != nil && ProjectName != "" {
		LoadProjectConfig(ProjectName)
	}
}

// Loads a TOML blob into the config map provided by the package.
func LoadBlob(blob string) (err error) {
	_, err = toml.Decode(blob, C)
}

// Loads a TOML file into the config map provided by the package.
func LoadFile(path string) (err error) {
	_, err = toml.DecodeFile(path, C)
}

// Loads an entire project's config (system level, user level, and local config) into the map provided by the package.
// Logs any errors that occur using the `github.com/SamWhited/logger` package.
// TODO: Don't just log errors now that this is a separate package; push them to a channel or something and let the user
// handle them.
func LoadProjectConfig(project string) {
	// Load system-level config
	err := LoadFile("/etc/" + projectName + "/config")
	if err != nil {
		logger.Info("No config found at `/etc/" + projectName + "/config'.")
	}

	// Load user-level config
	home, err := homedir.Dir()
	if err != nil {
		logger.Info("No home directory for current user; not loading user-level config")
	} else {
		err = LoadFile(home + "/.config/" + projectName + "/config")
		if err != nil {
			logger.Info("No config found at `" + home + "/.config/" + projectName + "/config'.")
		}
	}

	// Load local config
	err = LoadFile("config")
	if err != nil {
		logger.Info("No config found at `./config'.")
	}
}
