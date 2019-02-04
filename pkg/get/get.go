package get

import (
	"fmt"
	"strconv"

	"github.com/chengyumeng/khadijah/pkg/model"
	utillog "github.com/chengyumeng/khadijah/pkg/utils/log"
	"github.com/chengyumeng/khadijah/pkg/utils/table"
)

var (
	logger = utillog.NewAppLogger("pkg/get")
)

const ROW = "row"
const PRETTY = "pretty"

type GetProxy struct {
	Option Option
	table  table.Table
}

func NewProxy(opt Option) GetProxy {
	prx := GetProxy{
		Option: opt,
	}
	switch prx.Option.Output {
	case PRETTY:
		prx.table = table.NewTable(table.Horizontal)
	case ROW:
		prx.table = table.NewTable(table.Vertical)
	default:
		prx.table = table.NewTable(table.Horizontal)
	}
	return prx
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

func (g *GetProxy) print() {
	if g.table.IsEmpty() {
		logger.Warningln("There is no data in the table!")
	} else {
		g.table.Println()
	}
}

func (g *GetProxy) getNamespace() {
	if data := model.GetNamespaceBody(); data != nil {
		fmt.Printf("Name: %s Email:%s\n\n", data.Data.Name, data.Data.Email)
		g.table.SetHeaders([]string{"Id", "Name", "User", "CreateTime", "UpdateTime"})

		for _, v := range data.Data.Namespaces {
			if g.Option.NS == "" || g.Option.NS == v.Name {
				g.table.AddRow([]string{strconv.Itoa(int(v.Id)), v.Name, v.User, v.CreateTime.String(), v.UpdateTime.String()})
			}
		}
	}
	g.print()
}

func (g *GetProxy) getApp() {
	list := g.checkNS()

	g.table.SetHeaders([]string{"Id", "Name", "Namespace", "User", "CreateTime"})
	for _, ns := range list {
		if data := model.GetAppBody(ns.Id); data != nil {
			for _, v := range data.Data.Apps {
				if g.Option.App == "" || g.Option.App == v.Name {
					g.table.AddRow([]string{strconv.Itoa(int(v.Id)), v.Name, v.Namespace, v.User, v.CreateTime.String()})
				}
			}
		}

	}
	g.print()
}

func (g *GetProxy) GetPod(podType string) {
	list := g.checkNS()
	g.table.SetHeaders([]string{"Id", "Name", "Type", "APP", "Namespace", "User", "CreateTime"})
	for _, ns := range list {
		if app := model.GetAppBody(ns.Id); app != nil {
			for _, a := range app.Data.Apps {
				if g.Option.App == "" || g.Option.App == a.Name {
					if data := model.GetPodBody(a.Id, podType); data != nil {
						for _, pod := range data.Data.Pods {
							g.table.AddRow([]string{strconv.Itoa(int(pod.Id)), pod.Name, podType, pod.App.Name, pod.App.NSMetaData.Name, pod.User, pod.CreateTime.String()})
						}
					}
				}
			}
		}

	}
	g.print()
}

func (g *GetProxy) GetService() {
	list := g.checkNS()
	g.table.SetHeaders([]string{"Id", "Name", "Type", "APP", "Namespace", "User", "CreateTime"})
	for _, ns := range list {
		if app := model.GetAppBody(ns.Id); app != nil {
			for _, a := range app.Data.Apps {
				if g.Option.App == "" || g.Option.App == a.Name {
					if data := model.GetServiceBody(a.Id); data != nil {
						for _, svc := range data.Data.Services {
							g.table.AddRow([]string{strconv.Itoa(int(svc.Id)), svc.Name, model.ServiceType, a.Name, ns.Name, svc.User, svc.CreateTime.String()})
						}
					}
				}

			}
		}
	}
	g.print()
}

func (g *GetProxy) GetIngress() {
	list := g.checkNS()
	g.table.SetHeaders([]string{"Id", "Name", "Type", "APP", "Namespace", "User", "CreateTime"})
	for _, ns := range list {
		if app := model.GetAppBody(ns.Id); app != nil {
			for _, a := range app.Data.Apps {
				if g.Option.App == "" || g.Option.App == a.Name {
					if data := model.GetIngressBody(a.Id); data != nil {
						for _, ing := range data.Data.Ingresses {
							g.table.AddRow([]string{strconv.Itoa(int(ing.Id)), ing.Name, "Ingress", a.Name, ns.Name, ing.User, ing.CreateTime.String()})
						}
					}
				}

			}
		}
	}
	g.print()

}

func (g *GetProxy) GetAPIKey() {
	if data := model.GetAPIKeyBody(0); data != nil {
		g.table.SetHeaders([]string{"ID", "Name", "Type", "Resource ID", "User", "Description", "Token"})
		for _, api := range data.Data.APIkeys {
			g.table.AddRow([]string{strconv.Itoa(int(api.Id)), api.Name,
				strconv.Itoa(api.Type), strconv.Itoa(int(api.ResourceId)), api.User, api.Description, api.Token})
		}
	}
	g.print()
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
