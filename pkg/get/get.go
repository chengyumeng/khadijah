package get

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gosuri/uitable"
	"github.com/olekukonko/tablewriter"

	"github.com/chengyumeng/khadijah/pkg/model"
	utillog "github.com/chengyumeng/khadijah/pkg/utils/log"
)

var (
	logger = utillog.NewAppLogger("pkg/get")
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
	if g.Option.Option.Pod {
		g.GetPod(model.DeploymentType)
		g.GetPod(model.StatefulsetType)
		g.GetPod(model.DaemonsetType)
		g.GetPod(model.CronjobType)
	} else if g.Option.Option.Deployment {
		g.GetPod(model.DeploymentType)
	} else if g.Option.Option.DaemonSet {
		g.GetPod(model.DaemonsetType)
	} else if g.Option.Option.Statefulset {
		g.GetPod(model.StatefulsetType)
	} else if g.Option.Option.Cronjob {
		g.GetPod(model.CronjobType)
	} else if g.Option.Option.Service {
		g.GetService()
	} else if g.Option.Option.Ingress {
		g.GetIngress()
	} else if g.Option.Option.App {
		g.getApp()
	} else if g.Option.Option.Namespace {
		g.getNamespace()
	} else if g.Option.Option.APIkey {
		g.GetAPIKey()
	}
}

func (g *GetProxy) getNamespace() {
	data := model.GetNamespaceBody()
	if data == nil {
		return
	}
	fmt.Printf("Name: %s Email:%s\n\n", data.Data.Name, data.Data.Email)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "User", "CreateTime", "UpdateTime"})

	for _, v := range data.Data.Namespaces {
		if g.Option.NS == "" || g.Option.NS == v.Name {
			table.Append([]string{strconv.Itoa(int(v.Id)), v.Name, v.User, v.CreateTime.String(), v.UpdateTime.String()})
		}
	}
	table.Render()
}

func (g *GetProxy) getApp() {
	list := g.checkNS()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Namespace", "User", "CreateTime"})
	for _, ns := range list {
		data := model.GetAppBody(ns.Id)
		if data == nil {
			continue
		}

		for _, v := range data.Data.Apps {
			if g.Option.App == "" || g.Option.App == v.Name {
				table.Append([]string{strconv.Itoa(int(v.Id)), v.Name, v.Namespace, v.User, v.CreateTime.String()})
			}
		}

	}
	table.Render()
}

func (g *GetProxy) GetPod(podType string) {
	list := g.checkNS()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Type", "APP", "Namespace", "User", "CreateTime"})
	exist := false
	for _, ns := range list {
		if app := model.GetAppBody(ns.Id); app != nil {
			for _, a := range app.Data.Apps {
				if g.Option.App == "" || g.Option.App == a.Name {
					data := model.GetPodBody(a.Id, podType)
					for _, pod := range data.Data.Pods {
						exist = true
						table.Append([]string{strconv.Itoa(int(pod.Id)), pod.Name, podType, pod.App.Name, pod.App.NSMetaData.Name, pod.User, pod.CreateTime.String()})
					}
				}
			}
		}

	}
	if exist {
		table.Render()
	}
}

func (g *GetProxy) GetService() {
	list := g.checkNS()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Type", "APP", "Namespace", "User", "CreateTime"})
	exist := false
	for _, ns := range list {
		if app := model.GetAppBody(ns.Id); app != nil {
			for _, a := range app.Data.Apps {
				if g.Option.App == "" || g.Option.App == a.Name {
					data := model.GetServiceBody(a.Id)
					for _, svc := range data.Data.Services {
						exist = true
						table.Append([]string{strconv.Itoa(int(svc.Id)), svc.Name, model.ServiceType, a.Name, ns.Name, svc.User, svc.CreateTime.String()})
					}
				}

			}
		}
	}
	if exist {
		table.Render()
	}
}

func (g *GetProxy) GetIngress() {
	list := g.checkNS()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Type", "APP", "Namespace", "User", "CreateTime"})
	exist := false
	for _, ns := range list {
		if app := model.GetAppBody(ns.Id); app != nil {
			for _, a := range app.Data.Apps {
				if g.Option.App == "" || g.Option.App == a.Name {
					data := model.GetIngressBody(a.Id)
					for _, ing := range data.Data.Ingresses {
						exist = true
						table.Append([]string{strconv.Itoa(int(ing.Id)), ing.Name, "Ingress", a.Name, ns.Name, ing.User, ing.CreateTime.String()})
					}
				}

			}
		}
	}
	if exist {
		table.Render()
	}
}

func (g *GetProxy) GetAPIKey() {
	table := uitable.New()
	//table.MaxColWidth = 80
	table.Wrap = true // wrap columns

	data := model.GetAPIKeyBody(0)
	for _, api := range data.Data.APIkeys {
		table.AddRow("Id:", strconv.Itoa(int(api.Id)))
		table.AddRow("Name:", api.Name)
		table.AddRow("Type:", strconv.Itoa(api.Type))
		table.AddRow("Resource ID:", strconv.Itoa(int(api.ResourceId)))
		table.AddRow("User:", api.User)
		table.AddRow("Description:", api.Description)
		table.AddRow("Token:", api.Token)
		table.AddRow("") // blank
	}
	fmt.Println(table)
}

func (g *GetProxy) checkNS() (list []model.Namespace) {
	ns := model.GetNamespaceBody()
	if ns == nil {
		return
	}
	if g.Option.NS != "" {
		for _, n := range ns.Data.Namespaces {
			if n.Name == g.Option.NS {
				list = append(list, n)
			}
		}
	} else {
		for _, n := range ns.Data.Namespaces {
			list = append(list, n)
		}
	}
	if len(list) == 0 {
		logger.Error("Empty namespace list")
	}
	return
}
