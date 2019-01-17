package get

const (
	NamespaceType   string = "namespace"
	AppType         string = "app"
	DeploymentType  string = "deployment"
	DaemonsetType   string = "daemonset"
	CronjobType     string = "cronjob"
	StatefulsetType string = "statefulset"
	PodType         string = "pod"
	ServiceType     string = "service"
)

type Option struct {
	Resource  string
	Namespace string
	App       string
}


