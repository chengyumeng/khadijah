package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

type Option struct {
	Token  string `json:"token"`
	User   *User  `json:"user"`
	APIKey string `json:"apiKey"`
}

type User struct {
	Username string `json:"username"`
}

func LoadOption() (err error) {
	f, err := os.Open(path.Join(UserConfigDir, ConfigFile))
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &GlobalOption)
	return
}

func SetAPIKey(key string) (err error) {
	GlobalOption.APIKey = key
	return Save()
}

func SetUser(u *User) (err error) {
	GlobalOption.User = u
	return Save()
}

func SetToken(t string)(err error) {
	GlobalOption.Token = t
	return Save()
}
