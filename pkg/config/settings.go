package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

// Option is the basic option for khadijah-wayne
type Option struct {
	Token  string  `json:"token"`
	User   *User   `json:"user"`
	APIKey string  `json:"apiKey"`
	System *System `json:"system"`
}

// User is the user info of login user
type User struct {
	Username string `json:"username"`
}

// System is the basic system info for khadijah-wayne
type System struct {
	Version      string `json:"version"`
	BaseURL      string `json:"baseURL"`
	WebsocketURL string `json:"websocketURL"`
}

// LoadOption is the function to Load config option from local path
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

// SetAPIKey is the function to set APIKey to global option
func SetAPIKey(key string) (err error) {
	GlobalOption.APIKey = key
	return Save()
}

// SetUser is the function to set user to global option
func SetUser(u *User) (err error) {
	GlobalOption.User = u
	return Save()
}

// SetToken is the function to set token to global option
func SetToken(t string) (err error) {
	GlobalOption.Token = t
	return Save()
}

// SetSystem is the function to set system info to global option
func SetSystem(sys System) (err error) {
	GlobalOption.System = &sys
	return Save()
}
