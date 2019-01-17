package model

type AppBody struct {
	Data struct {
		Page
		Apps []App `json:"list"`
	} `json:"data"`
}

type NamespaceBody struct {
	Data struct {
		User
		Namespaces []Namespace `json:"namespaces"`
	} `json:"data"`
}



type PodBody struct {
	Data struct {
		Page
		Pods []Pod `json:"list"`
	} `json:"data"`
}



type ServiceBody struct {
	Data struct {
		Page
		Services []Service `json:"list"`
	} `json:"data"`
}