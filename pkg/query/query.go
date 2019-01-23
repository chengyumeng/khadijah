package query

import (
	"fmt"

	"github.com/chengyumeng/khadijah/pkg/model/openapi"
)

type QueryProxy struct {
}

func NewProxy() QueryProxy {
	return QueryProxy{}
}

func (g *QueryProxy) GetPodInfo(opt GetPodInfoOption) {
	data := openapi.Query("get_pod_info", []string{
		fmt.Sprintf("cluster=%s", opt.Cluster),
		fmt.Sprintf("labelSelector=%s", opt.LabelSelector),
	})
	fmt.Println(string(data))
}

func (g *QueryProxy) GetPodInfoFromIP(opt GetPodFromIPOption) {
	data := openapi.Query("get_pod_info_from_ip", []string{
		fmt.Sprintf("cluster=%s", opt.Cluster),
		fmt.Sprintf("ips=%s", opt.IP),
	})
	fmt.Println(string(data))
}

func (g *QueryProxy) GetResourceInfo(opt GetResourceInfoOption) {
	data := openapi.Query("get_resource_info", []string{
		fmt.Sprintf("type=%s", opt.Type),
		fmt.Sprintf("name=%s", opt.Name),
	})
	fmt.Println(string(data))
}

func (g *QueryProxy) GetDeploymentStatus(opt GetDeploymentStatusOption) {
	data := openapi.Query("get_deployment_status", []string{
		fmt.Sprintf("deployment=%s", opt.Deployment),
		fmt.Sprintf("cluster=%s", opt.Cluster),
		fmt.Sprintf("namespace=%s", opt.Namespace),
	})
	fmt.Println(string(data))
}