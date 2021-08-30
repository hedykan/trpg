package route

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func RouteInit() {
	http.HandleFunc("/", index)
	http.HandleFunc("/store/list", storyList)
	http.HandleFunc("/store/get", storyGet)
}

func res(data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = http.StatusOK
	res["data"] = data
	res["msg"] = "ok"
	return res
}

func get(r *http.Request) map[string]string {
	var res = make(map[string]string)
	keys := r.URL.Query()
	for index, value := range keys {
		res[index] = value[0]
	}

	return res
}
