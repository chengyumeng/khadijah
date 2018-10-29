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

type Page struct {
	PageNum int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
	TotalPage int64 `json:"totalPage"`
	TotalCount int64 `json:"totalCount"`
}

type AppBody struct {
	Data struct {
		Page
		Apps []App `json:"list"`
	}
}

type App struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Namespace string `json:"namespaceName"`
	User string `json:"user"`
	CreateTime *time.Time `json:"createTime"`
}
