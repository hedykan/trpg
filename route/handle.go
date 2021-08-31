package route

import (
	"fmt"
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
	resMiddle(w, r, nil)
}

func storyList(w http.ResponseWriter, r *http.Request) {
	resMiddle(w, r, con.StoryList())
}

func storyGet(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		panic(err)
	}
	resMiddle(w, r, con.StoryNodeGet(id))
}

func storyNodeAdd(w http.ResponseWriter, r *http.Request) {
	var query struct {
		val    string
		input  []int
		output []int
	}
	postJson(r, &query)
	fmt.Println("val:", query.input)
}

func runStatusReset(w http.ResponseWriter, r *http.Request) {
	con.RunStatusCreate()
	resMiddle(w, r, nil)
}

func runStatusList(w http.ResponseWriter, r *http.Request) {
	resMiddle(w, r, con.RunStatusList())
}

func runNowNodeGet(w http.ResponseWriter, r *http.Request) {
	resMiddle(w, r, con.RunNowNodeGet())
}

func runStep(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		panic(err)
	} else {
		con.RunStep(id)
	}
	resMiddle(w, r, nil)
}
