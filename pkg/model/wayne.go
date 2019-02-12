package model

import (
	"time"
)

// PageSize : Web 表单单页显示，不分页
const PageSize int = 1024 * 1024

// wayne 数据类型对照
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

// User : 用户个人信息
type User struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Display string `json:"display"`
	Admin   bool   `json:"admin"`
}

// Namespace : 命名空间基本数据结构
type Namespace struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Metadata   string     `json:"metadata"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	User       string     `json:"user"`
}

// Page : 页码基本数据结构
type Page struct {
	PageNum    int64 `json:"pageNo"`
	PageSize   int64 `json:"pageSize"`
	TotalPage  int64 `json:"totalPage"`
	TotalCount int64 `json:"totalCount"`
}

// App : 应用基本数据结构
type App struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Namespace  string     `json:"namespaceName"`
	User       string     `json:"user"`
	CreateTime *time.Time `json:"createTime"`
	NSMetaData Namespace  `json:"namespace"`
}

// Pod 基本数据结构
type Pod struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Metadata   string     `json:"metadata"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	User       string     `json:"user"`
	App        App        `json:"app"`
}

// Service : 负载均衡基本数据结构
type Service struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Metadata   string     `json:"metadata"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	User       string     `json:"user"`
	App        App        `json:"app"`
}

// Ingress 基本数据结构
type Ingress struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Metadata   string     `json:"metadata"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	User       string     `json:"user"`
	App        App        `json:"app"`
}

// APIKey 基本数据结构
type APIKey struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Token string `json:"token,omitempty"`
	// 0：全局 1：命名空间 2：项目
	Type        int        `json:"type"`
	ResourceID  int64      `json:"resourceId,omitempty"`
	Group       *Group     `json:"group,omitempty"`
	Description string     `json:"description,omitempty"`
	User        string     `json:"user,omitempty"`
	ExpireIn    int64      `json:"expireIn"`          // 过期时间，单位：秒
	Deleted     bool       `json:"deleted,omitempty"` // 是否生效
	CreateTime  *time.Time `json:"createTime,omitempty"`
	UpdateTime  *time.Time `json:"updateTime,omitempty"`
}

// Group 用户组基本数据结构
type Group struct {
	ID         int64      `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Comment    string     `json:"comment,omitempty"`
	Type       int        `json:"type"`
	CreateTime *time.Time `json:"createTime,omitempty"`
	UpdateTime *time.Time `json:"updateTime,omitempty"`
}

// Metadata 个性化数据
type Metadata struct {
	Namespace   string                 `json:"namespace"`
	Clusters    []string               `json:"clusters"`
	ClusterMeta map[string]interface{} `json:"clusterMeta"`
}
