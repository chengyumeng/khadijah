package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chengyumeng/khadijah/pkg/config"
)

// Get wayne namespace information
func GetNamespaceBody() *NamespaceBody {
	url := fmt.Sprintf("%s/%s", config.GlobalOption.System.BaseURL, "currentuser")
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", "Bearer "+config.GlobalOption.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Errorln(err)
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Warning(err)
	}
	if res.StatusCode != http.StatusOK {
		logger.Errorln(string(body))
		return nil
	}
	data := new(NamespaceBody)
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.Warning(err)
	}
	return data
}
