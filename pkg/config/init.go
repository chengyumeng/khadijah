package config

import (
	"os"
	"os/user"
	"path"
)

// khadijah global config option
var GlobalOption *Option = new(Option)

// user config directory, default be ""
var UserConfigDir string

// user config file name
var ConfigFile string = "config"

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
