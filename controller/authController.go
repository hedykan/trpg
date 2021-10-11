package controller

var CheckToken string

func AuthInit() {
	CheckToken = "13570890160"
}

func AuthCheck(token string) bool {
	return token == CheckToken
}
