package exec

type Option struct {
	Cluster    string
	Namespace  string
	Deployment string
	Pod        string
	Container  string
	Terminal   bool

	Cmd string
}
