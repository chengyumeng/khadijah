package describe

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ghodss/yaml"

	"github.com/chengyumeng/khadijah/pkg/model"
	"github.com/chengyumeng/khadijah/pkg/model/kubernetes"
	"github.com/chengyumeng/khadijah/pkg/utils/hash"
	utillog "github.com/chengyumeng/khadijah/pkg/utils/log"
	"github.com/chengyumeng/khadijah/pkg/utils/stringobj"
	"github.com/chengyumeng/khadijah/pkg/utils/table"
)

const (
	YAML   = "yaml"   // Print data with yaml type.
	JSON   = "json"   // Print data with json type.
	PRETTY = "pretty" // Print data with table type.
	ROW    = "row"    // Print data with row table type.
)

var (
	deploymentHeader = []string{"Name", "Namespace", "Cluster", "Labels", "Containers", "Replicas", "Message", "Pods"}
	serviceHeader    = []string{"Name", "Namespace", "Cluster", "Labels", "Type", "ClusterIP", "EXTERNAL-IP", "Ports", "SELECTOR"}
	ingressHeader    = []string{"Name", "Namespace", "Cluster", "Labels", "HOSTS"}
	configmapHeader  = []string{"Name", "Namespace", "Cluster", "Labels"}
	podHeader        = []string{"Name", "Namespace", "Cluster", "PodIP", "Node", "Restart Time", "Start Time"}

	logger = utillog.NewAppLogger("pkg/describe")
)

type DescribeProxy struct {
	Option Option
	table  table.Table
}

// Create a wayne describe proxy interface.
func NewProxy(opt Option) DescribeProxy {
	prx := DescribeProxy{
		Option: opt,
	}
	switch prx.Option.Output {
	case PRETTY:
		prx.table = table.NewTable(table.Horizontal)
	case ROW:
		prx.table = table.NewTable(table.Vertical)
	default:
		//
	}
	return prx
}

// Describe Kubernetes Object information.
func (g *DescribeProxy) Describe() {
	if g.Option.Option.Deployment != "" {
		g.Option.resource = model.DeploymentType
		g.showResourceState(g.Option.Option.Deployment)
	} else if g.Option.Option.Daemonset != "" {
		g.Option.resource = model.DaemonsetType
		g.showResourceState(g.Option.Option.Daemonset)
	} else if g.Option.Option.Statefulset != "" {
		g.Option.resource = model.StatefulsetType
		g.showResourceState(g.Option.Option.Statefulset)
	} else if g.Option.Option.Pod != "" {
		g.Option.resource = model.PodType
		g.showResourceState(g.Option.Option.Pod)
	} else if g.Option.Option.Service != "" {
		g.Option.resource = model.ServiceType
		g.showResourceState(g.Option.Option.Service)
	} else if g.Option.Option.Ingress != "" {
		g.Option.resource = model.IngressType
		g.showResourceState(g.Option.Option.Ingress)
	} else if g.Option.Option.Configmap != "" {
		g.Option.resource = model.ConfigmapType
		g.showResourceState(g.Option.Option.Configmap)
	} else if g.Option.Option.Pod != "" {
		g.Option.resource = model.PodType
		g.showResourceState(g.Option.Option.Pod)
	}
}

func (g *DescribeProxy) showResourceState(name string) {
	nslist := g.checkNS()
	for _, ns := range nslist {
		kns := new(model.Metadata)
		err := json.Unmarshal([]byte(ns.Metadata), &kns)
		if err != nil {
			logger.Errorln(err)
			return
		}
		if len(kns.Clusters) == 0 {
			if len(kns.ClusterMeta) > 0 {
				for k, _ := range kns.ClusterMeta {
					kns.Clusters = append(kns.Clusters, k)
				}
			} else if g.Option.Cluster == "" {
				logger.Warningln("You should insert cluster info!")
				return
			} else {
				kns.Clusters = append(kns.Clusters, g.Option.Cluster)
			}
		}
		for _, cluster := range kns.Clusters {
			if cluster == g.Option.Cluster || g.Option.Cluster == "" {
				data := kubernetes.GetResourceBody(name, int64(0), kns.Namespace, cluster, g.Option.resource, "")
				if data == nil || !hash.SetIfNotExist([]string{string(data)}) {
					continue
				}
				switch g.Option.Output {
				case YAML:
					data, err := yaml.JSONToYAML(data)
					if err != nil {
						logger.Errorln(err)
						return
					}
					fmt.Println(string(data))
				case JSON:
					fmt.Println(string(stringobj.String2Json(data)))
				case PRETTY, ROW:
					switch g.Option.resource {
					case model.DeploymentType, model.DaemonsetType, model.StatefulsetType:
						pods := kubernetes.ListPods(int64(0), kns.Namespace, cluster, "?"+g.Option.resource+"="+g.Option.Option.Deployment)
						arr := []string{}
						for _, p := range pods.Data {
							arr = append(arr, p.Name)
						}
						g.table.SetHeaders(deploymentHeader)
						if line := g.createDeploymentLine(data, cluster); len(line) > 0 {
							g.table.AddRow(append(line, strings.Join(arr, ",")))
						}
					case model.ServiceType:
						g.table.SetHeaders(serviceHeader)
						if line := g.createServiceLine(data, cluster); len(line) > 0 {
							g.table.AddRow(line)
						}
					case model.IngressType:
						g.table.SetHeaders(ingressHeader)
						if line := g.createIngressLine(data, cluster); len(line) > 0 {
							g.table.AddRow(line)
						}
					case model.ConfigmapType:
						g.table.SetHeaders(configmapHeader)
						if line := g.createConfigmapLine(data, cluster); len(line) > 0 {
							g.table.AddRow(line)
						}
					case model.PodType:
						g.table.SetHeaders(podHeader)
						pods := kubernetes.GetPod(int64(0), kns.Namespace, cluster, g.Option.Option.Pod)
						if line := g.createPodLine(pods.Data, cluster); len(line) > 0 {
							g.table.AddRow(line)
						}
					default:
						logger.Warningln(g.Option.resource)
					}
				}
			}
		}
	}
	if g.Option.Output == PRETTY || g.Option.Output == ROW {
		g.print()
	}
}

func (g *DescribeProxy) print() {
	if g.table.IsEmpty() {
		logger.Warningln("There is no data in the table!")
	} else {
		g.table.Println()
	}
}

func (g *DescribeProxy) createPodLine(pod *kubernetes.Pod, cluster string) []string {
	status := []string{}
	for _, s := range pod.ContainerStatus {
		status = append(status, fmt.Sprintf("%s:%d", s.Name, s.RestartCount))
	}
	return []string{
		pod.Name, pod.Namespace, cluster, pod.PodIp, pod.NodeName, strings.Join(status, ","), pod.StartTime.String(),
	}
}

func (g *DescribeProxy) createDeploymentLine(data []byte, cluster string) []string {
	obj := new(kubernetes.DeploymentBody)
	err := json.Unmarshal(data, &obj)
	if err != nil {
		logger.Errorln(err)
		return []string{}
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
		logger.Errorln(err)
		return []string{}
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

func (g *DescribeProxy) createIngressLine(data []byte, cluster string) []string {
	obj := new(kubernetes.IngressBody)
	err := json.Unmarshal(data, &obj)
	if err != nil {
		logger.Errorln(err)
		return []string{}
	}
	hosts := []string{}
	for _, r := range obj.Data.Spec.Rules {
		hosts = append(hosts, r.Host)
	}
	return []string{obj.Data.Name,
		obj.Data.Namespace, cluster,
		stringobj.Map2list(obj.Data.Labels), strings.Join(hosts, ",")}
}

func (g *DescribeProxy) createConfigmapLine(data []byte, cluster string) []string {
	obj := new(kubernetes.ConfigmapBody)
	err := json.Unmarshal(data, &obj)
	if err != nil {
		logger.Errorln(err)
		return []string{}
	}
	return []string{obj.Data.ObjectMeta.Name,
		obj.Data.ObjectMeta.Namespace, cluster,
		stringobj.Map2list(obj.Data.ObjectMeta.Labels)}
}

func (g *DescribeProxy) checkNS() (list []model.Namespace) {
	ns := model.GetNamespaceBody()
	if ns == nil {
		return
	}
	if g.Option.Namespace != "" {
		for _, n := range ns.Data.Namespaces {
			if n.Name == g.Option.Namespace {
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
