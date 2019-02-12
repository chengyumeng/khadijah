package query

// GetPodInfoOption : for wayne OpenAPI GetPodInfo
type GetPodInfoOption struct {
	LabelSelector string
	Cluster       string
}

// GetPodFromIPOption : for wayne OpenAPI GetPodFromIP
type GetPodFromIPOption struct {
	IP      string
	Cluster string
}

// GetResourceInfoOption : for wayne OpenAPI GetResourceInfo
type GetResourceInfoOption struct {
	Type string
	Name string
}

// GetDeploymentStatusOption : for wayne OpenAPI GetDeploymentStatus
type GetDeploymentStatusOption struct {
	Deployment string
	Namespace  string
	Cluster    string
}

// GetVIPInfoOption : for wayne OpenAPI GetVIPInfo
type GetVIPInfoOption struct {
	Port int
}
