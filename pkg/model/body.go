package model

// AppBody is the wayne APP http API response body
type AppBody struct {
	Data struct {
		Page
		Apps []App `json:"list"`
	} `json:"data"`
}

// NamespaceBody is the wayne Namespace http API response body
type NamespaceBody struct {
	Data struct {
		User
		Namespaces []Namespace `json:"namespaces"`
	} `json:"data"`
}

// PodBody is the wayne pod http API response body
type PodBody struct {
	Data struct {
		Page
		Pods []Pod `json:"list"`
	} `json:"data"`
}

// ServiceBody is the wayne service http API response body
type ServiceBody struct {
	Data struct {
		Page
		Services []Service `json:"list"`
	} `json:"data"`
}

// IngressBody is the wayne ingress http API response body
type IngressBody struct {
	Data struct {
		Page
		Ingresses []Ingress `json:"list"`
	} `json:"data"`
}

// APIKeyBody is the wayne APIKey http API response body
type APIKeyBody struct {
	Data struct {
		Page
		APIkeys []APIKey `json:"list"`
	}
}
