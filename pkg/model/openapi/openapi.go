package openapi

import (
	"fmt"
	"io/ioutil"
	"strings"
	"net/http"

	"github.com/chengyumeng/khadijah/pkg/config"
	"github.com/chengyumeng/khadijah/pkg/utils/log"
)

func Query(action string, params []string) []byte {
	url := fmt.Sprintf("%s/openapi/v1/gateway/action/%s?apikey=%s&%s", config.BaseURL, action, config.GlobalOption.APIKey, strings.Join(params, "&"))
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.AppLogger.Warning(err)
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println(string(body))
	}
	return body
}
