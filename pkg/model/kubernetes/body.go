package kubernetes

import (
	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
)

// DeploymentBody is wayne-kubernetes http API deployment response body
type DeploymentBody struct {
	Data struct {
		v1beta1.Deployment
	} `json:"data"`
}

// ServiceBody is wayne-kubernetes http API service response body
type ServiceBody struct {
	Data struct {
		v1.Service
	} `json:"data"`
}

// IngressBody is wayne-kubernetes http API ingress response body
type IngressBody struct {
	Data struct {
		extensionsv1beta1.Ingress
	} `json:"data"`
}

// ConfigmapBody is wayne-kubernetes http API configmap response body
type ConfigmapBody struct {
	Data struct {
		ConfigMap
	} `json:"data"`
}

// PodsBody is wayne-kubernetes http API pods response body
type PodsBody struct {
	Data []*Pod `json:"data"`
}

// PodBody is wayne-kubernetes http API pod response body
type PodBody struct {
	Data *Pod `json:"data"`
}
