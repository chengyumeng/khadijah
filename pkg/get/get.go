package get

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/chengyumeng/khadijah/pkg/config"
)

func Get(opt Option) {
	url := fmtURL(opt.Resource)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", "Bearer "+config.GlobalOption.Token)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println(string(body))
		return
	}
	data := new(NamespaceBody)
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

}

func fmtURL(resourceType string) string {
	switch resourceType {
	case NamespaceType:
		return fmt.Sprintf("%s/%s", config.BaseURL, "/currentuser")
	default:
		return ""
	}
}
