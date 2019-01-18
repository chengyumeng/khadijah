package kubernetes

import (
	"k8s.io/api/apps/v1beta1"
)

type DeploymentBody struct {
	Data struct {
		v1beta1.Deployment
	} `json:"data"`
}
