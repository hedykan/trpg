package route

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func RouteInit() {
	http.HandleFunc("/", index)
	http.HandleFunc("/story/list", storyList)
	http.HandleFunc("/story/get", storyGet)
	http.HandleFunc("/run/status_reset", runStatusReset)
	http.HandleFunc("/run/status_list", runStatusList)
	http.HandleFunc("/run/now_node_get", runNowNodeGet)
	http.HandleFunc("/run/step", runStep)
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

func resMiddle(w http.ResponseWriter, r *http.Request, oper interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	res := res(oper)
	enc.Encode(res)
}
