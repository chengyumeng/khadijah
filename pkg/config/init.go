package config

import (
	"os"
	"os/user"
	"path"
)

var (
	// GlobalOption is khadijah global config option
	GlobalOption = new(Option)
	// UserConfigDir is user config directory, default be ""
	UserConfigDir string
	// ConfigFile is user config file name
	ConfigFile = "config"
)

func init() {
	user, err := user.Current()
	if err != nil {
		logger.Error(err)
	}
	UserConfigDir = path.Join(user.HomeDir, ".khadijah")
	if _, err := os.Stat(UserConfigDir); os.IsNotExist(err) {
		err = os.Mkdir(UserConfigDir, 0744)
		if err != nil {
			logger.Error(err)
		}
	}
}
