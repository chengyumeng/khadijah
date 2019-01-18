package describe

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chengyumeng/khadijah/pkg/model"
	"github.com/chengyumeng/khadijah/pkg/model/kubernetes"
	"github.com/chengyumeng/khadijah/pkg/utils/log"
	"github.com/chengyumeng/khadijah/pkg/utils/stringobj"
	"github.com/ghodss/yaml"
	"github.com/olekukonko/tablewriter"
)

const YAML = "yaml"
const JSON = "json"
const PRETTY = "pretty"

type DescribeProxy struct {
	Option Option
}

func NewProxy(opt Option) DescribeProxy {
	return DescribeProxy{
		Option: opt,
	}
}

func (g *DescribeProxy) Describe() {
	if g.Option.Deployment != "" {
		g.Option.resource = model.DeploymentType
		g.showResourceState(g.Option.Deployment)
	} else if g.Option.Daemontset != "" {
		g.Option.resource = model.DaemonsetType
		g.showResourceState(g.Option.Daemontset)
	} else if g.Option.Statefulset != "" {
		g.Option.resource = model.StatefulsetType
		g.showResourceState(g.Option.Statefulset)
	}
}

func (g *DescribeProxy) showResourceState(name string) {
	data := model.GetNamespaceBody()
	nslist := []model.Namespace{}
	for _, ns := range data.Data.Namespaces {
		if ns.Name == g.Option.Namespace || g.Option.Namespace == "" {
			nslist = append(nslist, ns)
		}
	}
	tb := [][]string{}
	for _, ns := range nslist {
		kns := new(model.Metadata)
		err := json.Unmarshal([]byte(ns.Metadata), &kns)
		if err != nil {

		}
		for _, cluster := range kns.Clusters {
			if cluster == g.Option.Cluster || g.Option.Cluster == "" {
				data := kubernetes.GetResourceBody(name, int64(0), kns.Namespace, cluster, g.Option.resource)
				switch g.Option.Output {
				case YAML:
					data, err := yaml.JSONToYAML(data)
					if err != nil {

					}
					fmt.Println(string(data))
				case JSON:
					fmt.Println(string(data))
				case PRETTY:
					obj := new(kubernetes.DeploymentBody)
					err := json.Unmarshal(data, &obj)
					if err != nil {
						log.AppLogger.Error(err)
					}
					ic := make(map[string]string)
					for _, c := range obj.Data.Spec.Template.Spec.Containers {
						ic[c.Name] = c.Image
					}
					rc := fmt.Sprintf("%d/%d", obj.Data.Status.AvailableReplicas, obj.Data.Status.Replicas)
					msg := make(map[string]string)
					for _, c := range obj.Data.Status.Conditions {
						msg[c.LastUpdateTime.Local().String()] = c.Message
					}
					tb = append(tb, []string{obj.Data.Name, obj.Data.Namespace, cluster, stringobj.Map2list(obj.Data.Labels), stringobj.Map2list(ic), rc, stringobj.Map2list(msg)})
				}
			}
		}
	}
	if len(tb) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Namespace", "Cluster", "Labels", "Containers", "Replicas", "Message"})
		table.SetRowLine(true)
		table.SetRowSeparator("-")
		table.AppendBulk(tb)
		table.Render()
	}
}
