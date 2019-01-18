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
	"strings"
)

const YAML = "yaml"
const JSON = "json"
const PRETTY = "pretty"

var (
	DeploymentHeader = []string{"Name", "Namespace", "Cluster", "Labels", "Containers", "Replicas", "Message"}
	ServiceHeader    = []string{"Name", "Namespace", "Cluster", "Labels", "Type", "ClusterIP", "EXTERNAL-IP", "Ports", "SELECTOR"}
)

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
	} else if g.Option.Service != "" {
		g.Option.resource = model.ServiceType
		g.showResourceState(g.Option.Service)
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
	var header []string
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
					switch g.Option.resource {
					case model.DeploymentType:
					case model.DaemonsetType:
					case model.StatefulsetType:
						tb = append(tb, g.createDeploymentLine(data, cluster))
						header = DeploymentHeader
					case model.ServiceType:
						tb = append(tb, g.createServiceLine(data, cluster))
						header = ServiceHeader
					}
				}
			}
		}
	}
	if len(tb) > 0 {
		g.printTable(header, tb)
	}
}

func (g *DescribeProxy) printTable(header []string, lines [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetRowLine(true)
	table.SetRowSeparator("-")
	table.AppendBulk(lines)
	table.Render()
}

func (g *DescribeProxy) createDeploymentLine(data []byte, cluster string) []string {
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
	return []string{obj.Data.Name, obj.Data.Namespace, cluster, stringobj.Map2list(obj.Data.Labels), stringobj.Map2list(ic), rc, stringobj.Map2list(msg)}
}

func (g *DescribeProxy) createServiceLine(data []byte, cluster string) []string {
	obj := new(kubernetes.ServiceBody)
	err := json.Unmarshal(data, &obj)
	if err != nil {
		log.AppLogger.Error(err)
	}
	ps := []string{}
	for _, port := range obj.Data.Spec.Ports {
		ps = append(ps, fmt.Sprintf("%d:%d/%s", port.Port, port.TargetPort.IntVal, port.Protocol))
	}
	return []string{obj.Data.Name,
		obj.Data.Namespace, cluster,
		stringobj.Map2list(obj.Data.Labels),
		fmt.Sprintf("%v", obj.Data.Spec.Type),
		obj.Data.Spec.ClusterIP, strings.Join(obj.Data.Spec.ExternalIPs, ","),
		strings.Join(ps, ","), stringobj.Map2list(obj.Data.Spec.Selector)}
}
