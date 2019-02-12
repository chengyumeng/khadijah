package exec

// Option is the option of exec proxy
type Option struct {
	Cluster   string
	Namespace string
	Pod       string
	Container string
	Terminal  bool

	Cmd string
}
