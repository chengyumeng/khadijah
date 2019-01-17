package model

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"

	"github.com/chengyumeng/khadijah/pkg/utils/log"
	"github.com/chengyumeng/khadijah/pkg/config"
)


func GetNamespaceBody() *NamespaceBody {
	url := fmt.Sprintf("%s/%s", config.BaseURL, "currentuser")
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", "Bearer "+config.GlobalOption.Token)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.AppLogger.Warning(err)
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println(string(body))
		return nil
	}
	data := new(NamespaceBody)
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.AppLogger.Warning(err)
	}
	return data
}