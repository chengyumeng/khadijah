package config

import (
	"encoding/json"
	"os"
	"path"

	utillog "github.com/chengyumeng/khadijah/pkg/utils/log"
)

var (
	logger = utillog.NewAppLogger("pkg/describe")
)

// Save config to user local home path
func Save() (err error) {
	f, err := os.Create(path.Join(UserConfigDir, ConfigFile))
	if err != nil {
		return err
	} else {
		defer f.Close()
		data, err := json.MarshalIndent(GlobalOption, "", "  ")
		if err != nil {
			return err
		}
		_, err = f.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}
