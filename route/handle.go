package route

import (
	"net/http"
	"strconv"

	con "github.com/trpg/controller"
)

func index(w http.ResponseWriter, r *http.Request) {
	con.Test()
	resInput(w, r, nil)
}

func storyInit(w http.ResponseWriter, r *http.Request) {
	con.StoryCreate()
	resInput(w, r, con.StoryLoad())
}

func storyLoad(w http.ResponseWriter, r *http.Request) {
	resInput(w, r, con.StoryLoad())
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
	var query con.StoryNode
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

func storyNodeEdit(w http.ResponseWriter, r *http.Request) {
	var query con.StoryNode
	postJson(r, &query)
	ok := con.StoryNodeEdit(query.Id, query.Val, query.Input, query.Output)
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

func storySelecterAdd(w http.ResponseWriter, r *http.Request) {
	var query struct {
		NodeId int
		LinkId int
		Val    string
	}
	postJson(r, &query)
	resInput(w, r, con.StorySelecterAdd(query.NodeId, query.LinkId, query.Val))
}

func storySelecterDelete(w http.ResponseWriter, r *http.Request) {
	var query struct {
		NodeId int
		LinkId int
	}
	postJson(r, &query)
	resInput(w, r, con.StorySelecterDelete(query.NodeId, query.LinkId))
}

func runStatusReset(w http.ResponseWriter, r *http.Request) {
	con.RunStatusCreate()
	resInput(w, r, nil)
}

func runStatusList(w http.ResponseWriter, r *http.Request) {
	resInput(w, r, con.RunStatusList())
}

func runStoryBackgroundGet(w http.ResponseWriter, r *http.Request) {
	resInput(w, r, con.RunStoryBackgroundGet())
}

func runNowNodeGet(w http.ResponseWriter, r *http.Request) {
	resInput(w, r, con.RunNowNodeGet())
}

func runNowVoteGet(w http.ResponseWriter, r *http.Request) {
	resInput(w, r, con.RunNowVoteGet())
}

func runNowRecordList(w http.ResponseWriter, r *http.Request) {
	resInput(w, r, con.RunNowRecordList())
}

func runVoteAdd(w http.ResponseWriter, r *http.Request) {
	token := getToken(r)
	query := get(r)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		panic(err)
	} else {
		con.RunVoteAdd(id, token)
	}
	resInput(w, r, nil)
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

func attrList(w http.ResponseWriter, r *http.Request) {
	resInput(w, r, con.AttrList())
}

func attrNodeGet(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.AttrNodeGet(id))
}

func attrNodeAdd(w http.ResponseWriter, r *http.Request) {
	var query con.AttrNode
	postJson(r, &query)
	con.AttrNodeAdd(query.Val, query.Num)
	resInput(w, r, nil)
}

func attrNodeEdit(w http.ResponseWriter, r *http.Request) {
	var query con.AttrNode
	postJson(r, &query)
	resInput(w, r, con.AttrNodeEdit(query.Id, query.Val, query.Num))
}

func attrNodeDelete(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.AttrNodeDelete(id))
}

func authCheck(w http.ResponseWriter, r *http.Request) {
	token := getToken(r)
	resInput(w, r, con.AuthCheck(token))
}
