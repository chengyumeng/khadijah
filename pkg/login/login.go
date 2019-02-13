package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chengyumeng/khadijah/pkg/config"
	utillog "github.com/chengyumeng/khadijah/pkg/utils/log"
)

var (
	logger = utillog.NewAppLogger("pkg/login")
)

// Login is the interface for user login on wayne
func Login(opt Option) (err error) {
	url := fmt.Sprintf("%s/login/db?username=%s&password=%s", config.GlobalOption.System.BaseURL, opt.Username, opt.Password)

	req, _ := http.NewRequest("POST", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		logger.Errorln(string(body))
		return
	}
	data := new(Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	if err := config.SetToken(data.Data.Token); err != nil {
		return err
	}
	logger.Infoln("Login Success!")
	err = config.SetUser(&config.User{opt.Username})
	return err
}

// Clear is the interface for user clear wayne login token
func Clear() error {
	if err := config.SetToken(""); err != nil {
		return err
	}
	logger.Infoln("Logout Success!")
	err := config.SetUser(&config.User{""})
	return err
}
