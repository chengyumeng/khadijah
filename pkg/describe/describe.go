package describe

import (
	"encoding/json"
	"fmt"

	"github.com/chengyumeng/khadijah/pkg/model"
	"github.com/chengyumeng/khadijah/pkg/model/kubernetes"
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
	switch g.Option.Resource {
	default:
		g.showDeploymentState()
	}
}

func (g *DescribeProxy) showDeploymentState() {
	data := model.GetNamespaceBody()
	nslist := []model.Namespace{}
	for _,ns := range data.Data.Namespaces {
		if ns.Name == g.Option.Namespace || g.Option.Namespace == "" {
			nslist = append(nslist, ns)
		}
	}
	for _,ns := range nslist {
		kns := new(model.Metadata)
		err := json.Unmarshal([]byte(ns.Metadata),&kns)
		if err != nil {

		}
		data := kubernetes.GetResourceBody("move-abcdefghijklmnopqrstuvwsyzabcdefghijklmnsahlkj" ,int64(0),kns.Namespace,kns.Clusters[0])
		fmt.Println(string(data))
	}

}


