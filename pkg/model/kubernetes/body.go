package kubernetes

import (
	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
)

type DeploymentBody struct {
	Data struct {
		v1beta1.Deployment
	} `json:"data"`
}

type ServiceBody struct {
	Data struct {
		v1.Service
	} `json:"data"`
}

type IngressBody struct {
	Data struct {
		extensionsv1beta1.Ingress
	} `json:"data"`
}

type ConfigmapBody struct {
	Data struct {
		ConfigMap
	} `json:"data"`
}

type PodBody struct {
	Data []*Pod `json:"data"`
}
