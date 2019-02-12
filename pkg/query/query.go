package query

import (
	"fmt"

	"github.com/chengyumeng/khadijah/pkg/model/openapi"
	"github.com/chengyumeng/khadijah/pkg/utils/stringobj"
)

// Proxy is one wayne query proxy interface
type Proxy struct {
}

// NewProxy is the interface to create a wayne query OpenAPI proxy.
func NewProxy() Proxy {
	return Proxy{}
}

// GetPodInfo is the interface to get pod info from wayne OpenAPI
func (g *Proxy) GetPodInfo(opt GetPodInfoOption) {
	data := openapi.Query("get_pod_info", []string{
		fmt.Sprintf("cluster=%s", opt.Cluster),
		fmt.Sprintf("labelSelector=%s", opt.LabelSelector),
	})
	fmt.Println(string(stringobj.String2Json(data)))
}

// GetPodInfoFromIP is the interface to get pod info by IP from wayne OpenAPI
func (g *Proxy) GetPodInfoFromIP(opt GetPodFromIPOption) {
	data := openapi.Query("get_pod_info_from_ip", []string{
		fmt.Sprintf("cluster=%s", opt.Cluster),
		fmt.Sprintf("ips=%s", opt.IP),
	})
	fmt.Println(string(stringobj.String2Json(data)))
}

// GetResourceInfo is the interface to get resource info from wayne OpenAPI
func (g *Proxy) GetResourceInfo(opt GetResourceInfoOption) {
	data := openapi.Query("get_resource_info", []string{
		fmt.Sprintf("type=%s", opt.Type),
		fmt.Sprintf("name=%s", opt.Name),
	})
	fmt.Println(string(stringobj.String2Json(data)))
}

// GetDeploymentStatus is the interface to get deployment status from wayne OpenAPI
func (g *Proxy) GetDeploymentStatus(opt GetDeploymentStatusOption) {
	data := openapi.Query("get_deployment_status", []string{
		fmt.Sprintf("deployment=%s", opt.Deployment),
		fmt.Sprintf("cluster=%s", opt.Cluster),
		fmt.Sprintf("namespace=%s", opt.Namespace),
	})
	fmt.Println(string(stringobj.String2Json(data)))
}

// GetVIPInfo is the interface to get vip info from wayne OpenAPI
func (g *Proxy) GetVIPInfo(opt GetVIPInfoOption) {
	data := openapi.Query("get_vip_info", []string{
		fmt.Sprintf("port=%d", opt.Port),
	})
	fmt.Println(string(stringobj.String2Json(data)))
}
