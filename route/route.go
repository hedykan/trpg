package route

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func RouteInit() {
	http.Handle("/", checkMiddleware(index))
	http.Handle("/story/init", middleware(storyInit))
	http.Handle("/story/load", middleware(storyLoad))
	http.Handle("/story/list", middleware(storyList))
	http.Handle("/story/get", middleware(storyGet))
	http.Handle("/story/node_add", middleware(storyNodeAdd))
	http.Handle("/story/node_link", middleware(storyNodeLink))
	http.Handle("/story/node_edit", middleware(storyNodeEdit))
	http.Handle("/story/node_delete", middleware(storyNodeDelete))
	http.Handle("/story/selecter_add", middleware(storySelecterAdd))
	http.Handle("/story/selecter_delete", middleware(storySelecterDelete))

	http.Handle("/run/status_reset", middleware(runStatusReset))
	http.Handle("/run/status_list", middleware(runStatusList))
	http.Handle("/run/now_node_get", middleware(runNowNodeGet))
	http.Handle("/run/now_record_list", middleware(runNowRecordList))
	http.Handle("/run/step", middleware(runStep))
	http.Handle("/run/return", middleware(runReturn))

	http.Handle("/attr/list", middleware(attrList))
	http.Handle("/attr/node_get", middleware(attrNodeGet))
	http.Handle("/attr/node_add", middleware(attrNodeAdd))
	http.Handle("/attr/node_edit", middleware(attrNodeEdit))
	http.Handle("/attr/node_delete", middleware(attrNodeDelete))

	http.Handle("/auth/check", middleware(authCheck))
}

func res(data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = http.StatusOK
	res["data"] = data
	res["msg"] = "ok"
	return res
}

func getToken(r *http.Request) string {
	token := r.Header["Authorization"][0]
	return token
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

func resInput(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	res := res(data)
	enc.Encode(res)
}
