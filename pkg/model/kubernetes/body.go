package kubernetes

import (
	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/core/v1"
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
