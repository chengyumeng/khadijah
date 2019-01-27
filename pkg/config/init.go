package config

import (
	"os"
	"os/user"
	"path"
)

var GlobalOption *Option = new(Option)

var UserConfigDir string = ""
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
