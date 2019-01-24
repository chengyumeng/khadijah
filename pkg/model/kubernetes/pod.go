package kubernetes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chengyumeng/khadijah/pkg/config"
	"github.com/chengyumeng/khadijah/pkg/utils/log"
)

func GetResourceBody(resource string, appId int64, namespace string, cluster string, resourceType string, params string) []byte {
	url := fmt.Sprintf("%s/api/v1/kubernetes/apps/%d/%ss/%s/namespaces/%s/clusters/%s%s", config.GlobalOption.System.BaseURL, appId, resourceType, resource, namespace, cluster, params)
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

func ListPods(appId int64, namespace string, cluster string, params string) (obj PodsBody) {
	url := fmt.Sprintf("%s/api/v1/kubernetes/apps/%d/pods/namespaces/%s/clusters/%s%s", config.GlobalOption.System.BaseURL, appId, namespace, cluster, params)
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
	err = json.Unmarshal(body, &obj)
	return obj
}

func GetPod(appId int64, namespace string, cluster string, pod string) (obj PodBody) {
	url := fmt.Sprintf("%s/api/v1/kubernetes/apps/%d/pods/%s/namespaces/%s/clusters/%s", config.GlobalOption.System.BaseURL, appId, pod, namespace, cluster)
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
	err = json.Unmarshal(body, &obj)
	return obj
}
