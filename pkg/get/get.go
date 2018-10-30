package get

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/chengyumeng/khadijah/pkg/config"
	"github.com/olekukonko/tablewriter"
)

const pageSize int = 1024*1024 // 单页显示，不分页

type GetProxy struct {
	Option Option
}

func NewProxy(opt Option) GetProxy {
	return GetProxy{
		Option: opt,
	}
}

func (g *GetProxy) Get() {
	switch g.Option.Resource {
	case NamespaceType:
		g.getNamespace()
	case AppType:
		g.getApp()
	case DeploymentType:
	case StatefulsetType:
	case DaemonsetType:
	case CronjobType:
		g.GetPod(g.Option.Resource)
	case PodType:
		g.GetPod(DeploymentType)
		g.GetPod(StatefulsetType)
		g.GetPod(DaemonsetType)
		g.GetPod(CronjobType)
	default:
		fmt.Printf("Hello world!")
	}
}

func (g *GetProxy) getNamespace() {
	data := g.getNamespaceBody()
	fmt.Printf("Name: %s Email:%s\n\n", data.Data.Name, data.Data.Email)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "User", "CreateTime", "UpdateTime"})

	for _, v := range data.Data.Namespaces {
		table.Append([]string{strconv.Itoa(int(v.Id)), v.User, v.Name, v.CreateTime.String(), v.UpdateTime.String()})
	}
	table.Render()
}

func (g *GetProxy) getNamespaceBody() *NamespaceBody {
	url := fmt.Sprintf("%s/%s", config.BaseURL, "/currentuser")
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
		return nil
	}
	data := new(NamespaceBody)
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func (g *GetProxy) getApp() {
	nsIds := []int64{}
	ns := g.getNamespaceBody()
	if g.Option.Namespace != "" {
		for _, n := range ns.Data.Namespaces {
			if n.Name == g.Option.Namespace {
				nsIds = append(nsIds, n.Id)
			}
		}
		if len(nsIds) == 0 {
			panic("ERROR")
		}
	} else {
		for _, n := range ns.Data.Namespaces {
			nsIds = append(nsIds, n.Id)
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Namespace", "User", "CreateTime"})
	for _,id := range nsIds {
		data := g.getAppBody(id)
		if data == nil {
			continue
		}


		for _, v := range data.Data.Apps {
			table.Append([]string{strconv.Itoa(int(v.Id)), v.Name, v.Namespace, v.User, v.CreateTime.String()})
		}


	}
	table.Render()
}

func (g *GetProxy) getAppBody(nsId int64) *AppBody {
	url := fmt.Sprintf("%s/%s/%d/%s?pageSize=%d", config.BaseURL, "api/v1/namespaces",nsId,"apps",pageSize)
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
		return nil
	}
	data := new(AppBody)
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func (g *GetProxy) GetPod(podType string) {
	nsIds := []int64{}
	ns := g.getNamespaceBody()
	if g.Option.Namespace != "" {
		for _, n := range ns.Data.Namespaces {
			if n.Name == g.Option.Namespace {
				nsIds = append(nsIds, n.Id)
			}
		}
		if len(nsIds) == 0 {
			panic("ERROR")
		}
	} else {
		for _, n := range ns.Data.Namespaces {
			nsIds = append(nsIds, n.Id)
		}
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name","Type","APP", "Namespace", "User", "CreateTime"})
	exist := false
	for _,nsId := range nsIds {
		if app := g.getAppBody(nsId); app != nil {
			for _,a := range app.Data.Apps {
				if g.Option.App == "" || g.Option.App == a.Name {
					data := g.getPodBody(a.Id,podType)
					for _,pod := range data.Data.Pods {
						exist = true
						table.Append([]string{strconv.Itoa(int(pod.Id)), pod.Name,podType, pod.App.Name,pod.App.NSMetaData.Name, pod.User, pod.CreateTime.String()})
					}
				}
			}
		}

	}
	if exist {
		table.Render()
	}
}

func (g *GetProxy) getPodBody(appId int64,podType string) *PodBody {
	url := fmt.Sprintf("%s/%s/%d/%ss?pageSize=%d", config.BaseURL, "api/v1/apps",appId,podType,pageSize)
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
		return nil
	}
	data := new(PodBody)
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return data
}
