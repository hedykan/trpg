package route

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func RouteInit() {
	http.HandleFunc("/", index)
	http.HandleFunc("/story/list", storyList)
	http.HandleFunc("/story/get", storyGet)
	http.HandleFunc("/story/node_add", storyNodeAdd)

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

func postForm(r *http.Request, query []string) map[string]interface{} {
	var res = make(map[string]interface{})
	for i := 0; i < len(query); i++ {
		res[query[i]] = r.PostFormValue(query[i])
	}
	return res
}

func postJson(r *http.Request, obj interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, obj)
	if err != nil {
		panic(err)
	}
}

func resMiddle(w http.ResponseWriter, r *http.Request, oper interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	res := res(oper)
	enc.Encode(res)
}
