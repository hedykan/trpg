package route

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

// 把中间件赋值给固定变量，可以避免重复设置
func RouteInit() {
	// kp
	mid := checkMiddleware
	// 故事
	http.Handle("/", mid(index))
	http.Handle("/story/init", mid(storyInit))
	http.Handle("/story/load", mid(storyLoad))
	http.Handle("/story/list", mid(storyList))
	http.Handle("/story/get", mid(storyGet))
	http.Handle("/story/node_add", mid(storyNodeAdd))
	http.Handle("/story/node_edit", mid(storyNodeEdit))
	http.Handle("/story/node_delete", mid(storyNodeDelete))
	http.Handle("/story/selecter_add", mid(storySelecterAdd))
	http.Handle("/story/selecter_delete", mid(storySelecterDelete))

	// 跑团操作
	http.Handle("/run/status_reset", mid(runStatusReset))
	http.Handle("/run/step", mid(runStep))
	http.Handle("/run/return", mid(runReturn))

	// 属性操作
	http.Handle("/attr/node_add", mid(attrNodeAdd))
	http.Handle("/attr/node_edit", mid(attrNodeEdit))
	http.Handle("/attr/node_delete", mid(attrNodeDelete))

	// pc
	mid = middleware
	// 身份确认
	http.Handle("/auth/check", mid(authCheck))
	http.Handle("/auth/status", mid(authStatus))

	// 跑团操作
	http.Handle("/run/status_list", mid(runStatusList))
	http.Handle("/run/story_background_get", mid(runStoryBackgroundGet))
	http.Handle("/run/now_node_get", mid(runNowNodeGet))
	http.Handle("/run/now_vote_get", mid(runNowVoteGet))
	http.Handle("/run/vote_add", mid(runVoteAdd))
	http.Handle("/run/now_record_list", mid(runNowRecordList))

	// 属性操作
	http.Handle("/attr/list", mid(attrList))
	http.Handle("/attr/node_get", mid(attrNodeGet))
}

func res(data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = http.StatusOK
	res["data"] = data
	res["msg"] = "ok"
	return res
}

func errRes(code int) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = code
	switch code {
	case 201:
		res["msg"] = "权限不足"
		break
	}
	return res
}

func getToken(r *http.Request) string {
	token, ok := r.Header["Authorization"]
	if ok {
		return token[0]
	} else {
		return ""
	}
}

func getIp(r *http.Request) string {
	ip := r.Header.Get("Remote-Addr")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
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

func errResInput(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	res := errRes(code)
	enc.Encode(res)
}
