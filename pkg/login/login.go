package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/chengyumeng/khadijah/pkg/config"
)

func Login(opt Option) {
	url := fmt.Sprintf("%s/login/db?username=%s&password=%s", config.BaseURL, opt.Username, opt.Password)

	req, _ := http.NewRequest("POST", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println(string(body))
		return
	}
	data := new(Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	if err := save(data.Data.Token); err == nil {
		fmt.Println("Login Sucess!")
	} else {
		panic(err)
	}
}

func save(token string) error {
	f, err := os.Create(path.Join(config.UserConfigDir, config.TokenFile))
	if err != nil {
		return err
	} else {
		defer f.Close()
		f.WriteString(token)
	}
	return nil
}

func Clear() error {
	f, err := os.Create(path.Join(config.UserConfigDir, config.TokenFile))
	if err != nil {
		return err
	} else {
		defer f.Close()
	}
	return nil
}
