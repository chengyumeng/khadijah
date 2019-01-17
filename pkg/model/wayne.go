package model

import (
	"time"
)

const PageSize int = 1024 * 1024 // 单页显示，不分页

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



type Metadata struct {
	Namespace string `json:"namespace"`
	Clusters []string `json:"clusters"`
}
