package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chengyumeng/khadijah/pkg/config"
	"github.com/chengyumeng/khadijah/pkg/utils/log"
)

func Login(opt Option) (err error){
	url := fmt.Sprintf("%s/login/db?username=%s&password=%s", config.BaseURL, opt.Username, opt.Password)

	req, _ := http.NewRequest("POST", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		log.AppLogger.Error(string(body))
		return
	}
	data := new(Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	if err := config.SetToken(data.Data.Token); err == nil {
		log.AppLogger.Println("Login Success!")
	} else {
		return err
	}
	if err := config.SetUser( &config.User{opt.Username}); err != nil {
		return err
	}
	return nil
}

func Clear() error {
	if err := config.SetToken(""); err == nil {
		log.AppLogger.Println("Logout Success!")
	} else {
		return err
	}
	if err := config.SetUser( &config.User{ ""}); err != nil {
		return err
	}
	return nil
}
