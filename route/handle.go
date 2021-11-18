package route

import (
	"fmt"
	"net/http"
	"strconv"

	con "github.com/trpg/controller"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(getIp(r))
	con.Test()
	resInput(w, r, nil)
}

// func storyInit(w http.ResponseWriter, r *http.Request) {
// 	con.StoryCreate()
// 	// resInput(w, r, con.StoryLoad(con.StoryNodeArr, con.Addr))
// }

// func storyLoad(w http.ResponseWriter, r *http.Request) {
// 	// resInput(w, r, con.StoryLoad(con.StoryNodeArr, con.Addr))
// }

func storyList(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	id, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.RoomStoryList(id))
}

func storyGet(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	nodeid, err := strconv.Atoi(query["nodeId"])
	if err != nil {
		panic(err)
	}
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.RoomStoryNodeGet(roomId, nodeid))
}

func storyNodeAdd(w http.ResponseWriter, r *http.Request) {
	var query struct {
		RoomId int
		Val    string
		Input  []con.StorySeleter
		Output []con.StorySeleter
	}
	postJson(r, &query)
	ok := con.RoomStoryNodeAdd(query.RoomId, query.Val, query.Input, query.Output)
	resInput(w, r, ok)
}

func storyNodeEdit(w http.ResponseWriter, r *http.Request) {
	var query struct {
		RoomId int
		NodeId int
		Val    string
		Input  []con.StorySeleter
		Output []con.StorySeleter
	}
	postJson(r, &query)
	ok := con.RoomStoryNodeEdit(query.RoomId, query.NodeId, query.Val, query.Input, query.Output)
	resInput(w, r, ok)
}

func storyNodeDelete(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	nodeId, err := strconv.Atoi(query["nodeId"])
	if err != nil {
		panic(err)
	}
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	con.RoomStoryNodeDelete(roomId, nodeId)
	resInput(w, r, nil)
}

func storySelecterAdd(w http.ResponseWriter, r *http.Request) {
	var query struct {
		RoomId int
		NodeId int
		LinkId int
		Val    string
	}
	postJson(r, &query)

	resInput(w, r, con.RoomStorySelecterAdd(query.RoomId, query.NodeId, query.LinkId, query.Val))
}

func storySelecterDelete(w http.ResponseWriter, r *http.Request) {
	var query struct {
		RoomId int
		NodeId int
		LinkId int
	}
	postJson(r, &query)
	resInput(w, r, con.RoomStorySelecterDelete(query.RoomId, query.NodeId, query.LinkId))
}

// func runStatusReset(w http.ResponseWriter, r *http.Request) {
// 	// con.RunStatusCreate()
// 	resInput(w, r, nil)
// }

func runStatusList(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.RoomRunStatusList(roomId))
}

func runStoryBackgroundGet(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.RoomRunBackgroundGet(roomId))
}

func runNowNodeGet(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.RoomRunNowNodeGet(roomId))
}

func runNowVoteGet(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.RoomRunNowVoteGet(roomId))
}

func runNowRecordList(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.RoomRunNowRecordList(roomId))
}

func runVoteAdd(w http.ResponseWriter, r *http.Request) {
	// token := getToken(r)
	token := getIp(r)
	query := get(r)
	nodeId, err := strconv.Atoi(query["nodeId"])
	if err != nil {
		panic(err)
	}
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.RoomRunVoteAdd(roomId, nodeId, token))
}

func runStep(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	nodeId, err := strconv.Atoi(query["nodeId"])
	if err != nil {
		panic(err)
	}
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	con.RoomRunStep(roomId, nodeId)
	resInput(w, r, nil)
}

func runReturn(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	nodeId, err := strconv.Atoi(query["nodeId"])
	if err != nil {
		panic(err)
	}
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	con.RoomRunReturn(roomId, nodeId)
	resInput(w, r, nil)
}

func attrList(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.RoomAttrList(roomId))
}

func attrNodeGet(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	attrId, err := strconv.Atoi(query["attrId"])
	if err != nil {
		panic(err)
	}
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.RoomAttrNodeGet(roomId, attrId))
}

func attrNodeAdd(w http.ResponseWriter, r *http.Request) {
	var query struct {
		RoomId int
		Val    string
		Num    int
	}
	postJson(r, &query)
	con.RoomAttrNodeAdd(query.RoomId, query.Val, query.Num)
	resInput(w, r, nil)
}

func attrNodeEdit(w http.ResponseWriter, r *http.Request) {
	var query struct {
		RoomId int
		AttrId int
		Val    string
		Num    int
	}
	postJson(r, &query)
	resInput(w, r, con.RoomAttrNodeEdit(query.RoomId, query.AttrId, query.Val, query.Num))
}

func attrNodeDelete(w http.ResponseWriter, r *http.Request) {
	query := get(r)
	attrId, err := strconv.Atoi(query["attrId"])
	if err != nil {
		panic(err)
	}
	roomId, err := strconv.Atoi(query["roomId"])
	if err != nil {
		panic(err)
	}
	resInput(w, r, con.RoomAttrNodeDelete(roomId, attrId))
}

func roomList(w http.ResponseWriter, r *http.Request) {
	resInput(w, r, con.RoomList())
}

func authCheck(w http.ResponseWriter, r *http.Request) {
	token := getToken(r)
	resInput(w, r, con.AuthCheck(token, "kp", 1))
}

func authStatus(w http.ResponseWriter, r *http.Request) {
	token := getToken(r)
	resInput(w, r, con.AuthStatus(token))
}
