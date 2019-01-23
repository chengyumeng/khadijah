package exec

type Option struct {
	Cluster   string
	Namespace string
	Pod       string
	Container string
	Terminal  bool

	Cmd string
}
