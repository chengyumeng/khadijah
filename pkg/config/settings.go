package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

type Option struct {
	Token  string  `json:"token"`
	User   *User   `json:"user"`
	APIKey string  `json:"apiKey"`
	System *System `json:"system"`
}

type User struct {
	Username string `json:"username"`
}

type System struct {
	Version string `json:"version"`
	BaseURL string `json:"baseURL"`
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

func SetToken(t string) (err error) {
	GlobalOption.Token = t
	return Save()
}

func SetSystem(sys System) (err error) {
	GlobalOption.System = &sys
	return Save()
}
