package login

type Option struct {
	Username string
	Password string
}

type Body struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}
