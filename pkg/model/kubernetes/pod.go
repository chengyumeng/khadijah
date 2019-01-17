package kubernetes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chengyumeng/khadijah/pkg/config"
	"github.com/chengyumeng/khadijah/pkg/utils/log"
)

func GetResourceBody(resource string, appId int64, namespace string, cluster string, resourceType string) []byte {
	url := fmt.Sprintf("%s/api/v1/kubernetes/apps/%d/%ss/%s/namespaces/%s/clusters/%s", config.BaseURL, appId, resourceType, resource, namespace, cluster)
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
