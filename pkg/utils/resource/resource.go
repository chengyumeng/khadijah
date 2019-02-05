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

type DescribeOption struct {
	Deployment  string
	Daemonset   string
	Statefulset string
	Service     string
	Ingress     string
	Configmap   string
	Pod         string
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

func ParserResource(args []string) *DescribeOption {
	if len(args) == 2 {
		switch args[0] {
		case "Deployment", "deployment", "deploy":
			return &DescribeOption{Deployment: args[1]}
		case "Daemonset", "daemonset":
			return &DescribeOption{Daemonset: args[1]}
		case "Statefulset", "statefulset":
			return &DescribeOption{Statefulset: args[1]}
		case "Service", "service", "svc":
			return &DescribeOption{Service: args[1]}
		case "Configmap", "configmap", "cfg", "config":
			return &DescribeOption{Configmap: args[1]}
		case "Ingress", "ingress", "ing":
			return &DescribeOption{Ingress: args[1]}
		case "Pod", "pod":
			return &DescribeOption{Pod: args[1]}
		default:
			return nil
		}
	}
	return nil
}
