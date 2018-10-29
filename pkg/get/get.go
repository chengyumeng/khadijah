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
	url := fmt.Sprintf("%s/%s/%d/%s?pageSize=10241024", config.BaseURL, "api/v1/namespaces",nsId,"apps")
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
