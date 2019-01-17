package describe

import (
	"encoding/json"
	"fmt"

	"github.com/chengyumeng/khadijah/pkg/model"
	"github.com/chengyumeng/khadijah/pkg/model/kubernetes"
	"github.com/ghodss/yaml"
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
					fmt.Println("init")
				}
			}
		}
	}

}
