package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chengyumeng/khadijah/pkg/config"
)

// GetPodBody is the interface to get pod body from wayne API
func GetPodBody(appID int64, podType string) *PodBody {
	url := fmt.Sprintf("%s/%s/%d/%ss?pageSize=%d", config.GlobalOption.System.BaseURL, "api/v1/apps", appID, podType, PageSize)
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
		logger.Errorln(err)
		return nil
	}
	if res.StatusCode != http.StatusOK {
		logger.Println(string(body))
		return nil
	}
	data := new(PodBody)
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.Errorln(err)
		return nil
	}
	return data
}
