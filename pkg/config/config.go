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
		f.Write(data)
	}
	return nil
}
