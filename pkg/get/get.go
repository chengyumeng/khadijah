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

// Print data with table/row table type.
const (
	ROW    = "row"
	PRETTY = "pretty"
)

// Proxy is one wayne get proxy interface
type Proxy struct {
	Option Option
	table  table.Table
}

// NewProxy is the interface to create a wayne get proxy.
func NewProxy(opt Option) Proxy {
	prx := Proxy{
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

// Get wayne Object information.
func (g *Proxy) Get() {
	if g.Option.Option.Pod {
		g.getPod(model.DeploymentType)
		g.getPod(model.StatefulsetType)
		g.getPod(model.DaemonsetType)
		g.getPod(model.CronjobType)
	} else if g.Option.Option.Deployment {
		g.getPod(model.DeploymentType)
	} else if g.Option.Option.DaemonSet {
		g.getPod(model.DaemonsetType)
	} else if g.Option.Option.Statefulset {
		g.getPod(model.StatefulsetType)
	} else if g.Option.Option.Cronjob {
		g.getPod(model.CronjobType)
	} else if g.Option.Option.Service {
		g.getService()
	} else if g.Option.Option.Ingress {
		g.getIngress()
	} else if g.Option.Option.App {
		g.getApp()
	} else if g.Option.Option.Namespace {
		g.getNamespace()
	} else if g.Option.Option.APIkey {
		g.getAPIKey()
	}
}

func (g *Proxy) print() {
	if g.table.IsEmpty() {
		logger.Warningln("There is no data in the table!")
	} else {
		g.table.Println()
	}
}

func (g *Proxy) getNamespace() {
	if data := model.GetNamespaceBody(); data != nil {
		fmt.Printf("Name: %s Email:%s\n\n", data.Data.Name, data.Data.Email)
		g.table.SetHeaders([]string{"Id", "Name", "User", "CreateTime", "UpdateTime"})

		for _, v := range data.Data.Namespaces {
			if g.Option.NS == "" || g.Option.NS == v.Name {
				g.table.AddRow([]string{strconv.Itoa(int(v.ID)), v.Name, v.User, v.CreateTime.String(), v.UpdateTime.String()})
			}
		}
	}
	g.print()
}

func (g *Proxy) getApp() {
	list := g.checkNS()

	g.table.SetHeaders([]string{"Id", "Name", "Namespace", "User", "CreateTime"})
	for _, ns := range list {
		if data := model.GetAppBody(ns.ID); data != nil {
			for _, v := range data.Data.Apps {
				if g.Option.App == "" || g.Option.App == v.Name {
					g.table.AddRow([]string{strconv.Itoa(int(v.ID)), v.Name, v.Namespace, v.User, v.CreateTime.String()})
				}
			}
		}

	}
	g.print()
}

func (g *Proxy) getPod(podType string) {
	list := g.checkNS()
	g.table.SetHeaders([]string{"Id", "Name", "Type", "APP", "Namespace", "User", "CreateTime"})
	for _, ns := range list {
		if app := model.GetAppBody(ns.ID); app != nil {
			for _, a := range app.Data.Apps {
				if g.Option.App == "" || g.Option.App == a.Name {
					if data := model.GetPodBody(a.ID, podType); data != nil {
						for _, pod := range data.Data.Pods {
							g.table.AddRow([]string{strconv.Itoa(int(pod.ID)), pod.Name, podType, pod.App.Name, pod.App.NSMetaData.Name, pod.User, pod.CreateTime.String()})
						}
					}
				}
			}
		}

	}
	g.print()
}

func (g *Proxy) getService() {
	list := g.checkNS()
	g.table.SetHeaders([]string{"Id", "Name", "Type", "APP", "Namespace", "User", "CreateTime"})
	for _, ns := range list {
		if app := model.GetAppBody(ns.ID); app != nil {
			for _, a := range app.Data.Apps {
				if g.Option.App == "" || g.Option.App == a.Name {
					if data := model.GetServiceBody(a.ID); data != nil {
						for _, svc := range data.Data.Services {
							g.table.AddRow([]string{strconv.Itoa(int(svc.ID)), svc.Name, model.ServiceType, a.Name, ns.Name, svc.User, svc.CreateTime.String()})
						}
					}
				}

			}
		}
	}
	g.print()
}

func (g *Proxy) getIngress() {
	list := g.checkNS()
	g.table.SetHeaders([]string{"Id", "Name", "Type", "APP", "Namespace", "User", "CreateTime"})
	for _, ns := range list {
		if app := model.GetAppBody(ns.ID); app != nil {
			for _, a := range app.Data.Apps {
				if g.Option.App == "" || g.Option.App == a.Name {
					if data := model.GetIngressBody(a.ID); data != nil {
						for _, ing := range data.Data.Ingresses {
							g.table.AddRow([]string{strconv.Itoa(int(ing.ID)), ing.Name, "Ingress", a.Name, ns.Name, ing.User, ing.CreateTime.String()})
						}
					}
				}

			}
		}
	}
	g.print()

}

func (g *Proxy) getAPIKey() {
	if data := model.GetAPIKeyBody(0); data != nil {
		g.table.SetHeaders([]string{"ID", "Name", "Type", "Resource ID", "User", "Description", "Token"})
		for _, api := range data.Data.APIkeys {
			g.table.AddRow([]string{strconv.Itoa(int(api.ID)), api.Name,
				strconv.Itoa(api.Type), strconv.Itoa(int(api.ResourceID)), api.User, api.Description, api.Token})
		}
	}
	g.print()
}

func (g *Proxy) checkNS() (list []model.Namespace) {
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
		logger.Errorln("Empty namespace list")
	}
	return
}
