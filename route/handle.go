package route

import (
	"net/http"
	"strconv"

	con "github.com/trpg/controller"
)

func index(w http.ResponseWriter, r *http.Request) {
	// con.StoryCreate()
	// con.RunStatusCreate()
	// con.StoryNodeAdd("选项1", []int{0}, []int{1})
	// con.StoryNodeAdd("选项2", []int{0}, []int{1})
	// con.StoryNodeLink("后续选项1", 2, 1)
	// con.StoryNodeLink("后续选项2", 3, 1)
	resInput(w, r, nil)
}

func storyInit(w http.ResponseWriter, r *http.Request) {
	con.StoryCreate()
	resInput(w, r, con.StoryInit())
}

func storyList(w http.ResponseWriter, r *http.Request) {
	resInput(w, r, con.StoryList())
}

func storyGet(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.StoryNodeGet(id))
}

func storyNodeAdd(w http.ResponseWriter, r *http.Request) {
	var query struct {
		Val    string
		Input  []con.StorySeleter
		Output []con.StorySeleter
	}
	postJson(r, &query)
	ok := con.StoryNodeAdd(query.Val, query.Input, query.Output)
	resInput(w, r, ok)
}

func storyNodeLink(w http.ResponseWriter, r *http.Request) {
	var query struct {
		Val    string
		Input  con.StorySeleter
		Output con.StorySeleter
	}
	postJson(r, &query)
	ok := con.StoryNodeLink(query.Val, query.Input, query.Output)
	resInput(w, r, ok)
}

func storyNodeDelete(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		panic(err)
	} else {
		con.StoryNodeDelete(id)
	}
	resInput(w, r, nil)
}

func runStatusReset(w http.ResponseWriter, r *http.Request) {
	con.RunStatusCreate()
	resInput(w, r, nil)
}

func runStatusList(w http.ResponseWriter, r *http.Request) {
	resInput(w, r, con.RunStatusList())
}

func runNowNodeGet(w http.ResponseWriter, r *http.Request) {
	resInput(w, r, con.RunNowNodeGet())
}

func runNowRecordList(w http.ResponseWriter, r *http.Request) {
	resInput(w, r, con.RunNowRecordList())
}

func runStep(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		panic(err)
	} else {
		con.RunStep(id)
	}
	resInput(w, r, nil)
}

func runReturn(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		panic(err)
	}
	con.RunReturn(id)
	resInput(w, r, nil)
}
