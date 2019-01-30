package kubernetes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chengyumeng/khadijah/pkg/config"
	utillog "github.com/chengyumeng/khadijah/pkg/utils/log"
)

var (
	logger = utillog.NewAppLogger("pkg/model/kubernetes")
)

func GetResourceBody(resource string, appId int64, namespace string, cluster string, resourceType string, params string) []byte {
	url := fmt.Sprintf("%s/api/v1/kubernetes/apps/%d/%ss/%s/namespaces/%s/clusters/%s%s", config.GlobalOption.System.BaseURL, appId, resourceType, resource, namespace, cluster, params)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Errorln(err)
		return nil
	}
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
		logger.Warning(string(body))
	}
	return body
}

func ListPods(appId int64, namespace string, cluster string, params string) (obj PodsBody) {
	url := fmt.Sprintf("%s/api/v1/kubernetes/apps/%d/pods/namespaces/%s/clusters/%s%s", config.GlobalOption.System.BaseURL, appId, namespace, cluster, params)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Errorln(err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+config.GlobalOption.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Errorln(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Warning(err)
	}
	if res.StatusCode != http.StatusOK {
		logger.Warning(string(body))
	}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		logger.Errorln(err)
	}
	return obj
}

func GetPod(appId int64, namespace string, cluster string, pod string) (obj PodBody) {
	url := fmt.Sprintf("%s/api/v1/kubernetes/apps/%d/pods/%s/namespaces/%s/clusters/%s", config.GlobalOption.System.BaseURL, appId, pod, namespace, cluster)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Errorln(err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+config.GlobalOption.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Errorln(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Warning(err)
	}
	if res.StatusCode != http.StatusOK {
		logger.Warning(string(body))
	}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		logger.Errorln(err)
	}
	return obj
}
