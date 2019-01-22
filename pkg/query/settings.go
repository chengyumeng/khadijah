package query

type GetPodInfoOption struct {
	LabelSelector string
	Cluster       string
}

type GetPodFromIPOption struct {
	IP      string
	Cluster string
}

type GetResourceInfoOption struct {
	Type string
	Name string
}

type GetDeploymentStatusOption struct {
	Deployment string
	Namespace  string
	Cluster    string
}

type GetVIPInfo struct {
}
