package login

// 登录请求接口数据格式
type Option struct {
	Username string
	Password string
}

// Web 请求返回 body 格式
type Body struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}
