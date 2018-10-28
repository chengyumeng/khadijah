package config

import (
	"path"
	"os"
	"io/ioutil"
)

type Option struct {
	Token string
}

func LoadOption()(opt *Option,err error){
	f, err := os.Open(path.Join(UserConfigDir,TokenFile))
	if err != nil {
		return nil,err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil,err
	}
	opt.Token = string(data)
	return
}
