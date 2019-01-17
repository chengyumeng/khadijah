package kubernetes

import (
	"net/http"
	"fmt"
	"io/ioutil"

	"github.com/chengyumeng/khadijah/pkg/utils/log"
	"github.com/chengyumeng/khadijah/pkg/config"
)


func GetResourceBody(resource string,appId int64,namespace string,cluster string) []byte {
	url := fmt.Sprintf("%s/api/v1/kubernetes/apps/%d/deployments/%s/namespaces/%s/clusters/%s", config.BaseURL, appId,resource,namespace,cluster)
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
	}
	return body
}
