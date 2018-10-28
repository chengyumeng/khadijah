package config

import (
	"os/user"
	"os"
	"path"
)

var Version string = ""

var GlobalOption *Option

var BaseURL string = "http://open.qihoo.cloud"

var UserConfigDir string = ""
var TokenFile string = ".token"

func init() {
	user,err := user.Current()
	if err != nil {
		panic(err)
	}
	UserConfigDir = path.Join(user.HomeDir, ".khadijah")
	if _, err := os.Stat(UserConfigDir); os.IsNotExist(err) {
		err = os.Mkdir(UserConfigDir, 0744)
		if err != nil {
			panic(err)
		}
	}
}
