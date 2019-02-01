package get

type Option struct {
	resource    string
	Deployment  bool
	Daemonset   bool
	Statefulset bool
	Pod         bool
	Cronjob     bool
	Service     bool
	Ingress     bool
	Application bool
	Namespace   bool
	NS          string
	App         string
}
