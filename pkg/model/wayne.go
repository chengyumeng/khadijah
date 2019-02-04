package model

import (
	"time"
)

const PageSize int = 1024 * 1024 // 单页显示，不分页

const (
	NamespaceType   string = "namespace"
	AppType         string = "app"
	DeploymentType  string = "deployment"
	DaemonsetType   string = "daemonset"
	CronjobType     string = "cronjob"
	StatefulsetType string = "statefulset"
	PodType         string = "pod"
	ServiceType     string = "service"
	IngressType     string = "ingresse"
	ConfigmapType   string = "configmap"
)

type User struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Display string `json:"display"`
	Admin   bool   `json:"admin"`
}

type Namespace struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	Metadata   string     `json:"metadata"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	User       string     `json:"user"`
}

type Page struct {
	PageNum    int64 `json:"pageNo"`
	PageSize   int64 `json:"pageSize"`
	TotalPage  int64 `json:"totalPage"`
	TotalCount int64 `json:"totalCount"`
}

type App struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	Namespace  string     `json:"namespaceName"`
	User       string     `json:"user"`
	CreateTime *time.Time `json:"createTime"`
	NSMetaData Namespace  `json:"namespace"`
}

type Pod struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	Metadata   string     `json:"metadata"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	User       string     `json:"user"`
	App        App        `json:"app"`
}

type Service struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	Metadata   string     `json:"metadata"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	User       string     `json:"user"`
	App        App        `json:"app"`
}

type Ingress struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	Metadata   string     `json:"metadata"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	User       string     `json:"user"`
	App        App        `json:"app"`
}

type APIKey struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Token string `json:"token,omitempty"`
	// 0：全局 1：命名空间 2：项目
	Type        int        `json:"type"`
	ResourceId  int64      `json:"resourceId,omitempty"`
	Group       *Group     `json:"group,omitempty"`
	Description string     `json:"description,omitempty"`
	User        string     `json:"user,omitempty"`
	ExpireIn    int64      `json:"expireIn"`          // 过期时间，单位：秒
	Deleted     bool       `json:"deleted,omitempty"` // 是否生效
	CreateTime  *time.Time `json:"createTime,omitempty"`
	UpdateTime  *time.Time `json:"updateTime,omitempty"`
}

type Group struct {
	Id         int64      `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Comment    string     `json:"comment,omitempty"`
	Type       int        `json:"type"`
	CreateTime *time.Time `json:"createTime,omitempty"`
	UpdateTime *time.Time `json:"updateTime,omitempty"`
}

type Metadata struct {
	Namespace   string                 `json:"namespace"`
	Clusters    []string               `json:"clusters"`
	ClusterMeta map[string]interface{} `json:"clusterMeta"`
}
