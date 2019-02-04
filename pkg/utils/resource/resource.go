package resource

type Option struct {
	Namespace   bool
	App         bool
	Deployment  bool
	DaemonSet   bool
	Statefulset bool
	Service     bool
	Configmap   bool
	Ingress     bool
	Pod         bool
	Cronjob     bool
	APIkey      bool
}

func ParserArgs(args []string) *Option {
	if len(args) == 1 {
		switch args[0] {
		case "Namespace", "NS", "ns", "namespace":
			return &Option{Namespace: true}
		case "Application", "application", "app", "App":
			return &Option{App: true}
		case "Deployment", "deployment", "deploy":
			return &Option{Deployment: true}
		case "Daemonset", "daemonset":
			return &Option{DaemonSet: true}
		case "Statefulset", "statefulset":
			return &Option{Statefulset: true}
		case "Service", "service", "svc":
			return &Option{Service: true}
		case "Configmap", "configmap", "cfg", "config":
			return &Option{Configmap: true}
		case "Ingress", "ingress", "ing":
			return &Option{Ingress: true}
		case "Pod", "pod":
			return &Option{Pod: true}
		case "Cronjob", "cronjob":
			return &Option{Cronjob: true}
		case "APIkey", "apikey", "api", "API", "key", "Key", "k":
			return &Option{APIkey: true}
		default:
			return nil
		}
	}
	return nil
}
