package get

import "time"

const (
	NamespaceType  string = "namespace"
	AppType        string = "app"
	DeploymentType string = "deployment"
	ServiceType    string = "service"
)

type Option struct {
	Resource  string
	Namespace string
	App       string
}

type NamespaceBody struct {
	Data struct {
		User
		Namespaces []Namespace `json:"namespaces"`
	} `json:"data"`

}

type User struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Display string `json:"display"`
	Admin bool `json:"admin"`
}

type Namespace struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Metadata string `json:"metadata"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	User string `json:"user"`
}
