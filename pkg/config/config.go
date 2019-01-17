package config

import (
	"os"
	"path"
	"encoding/json"
)

func Save() (err error){
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
