package model

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Room struct {
	RoomId        int
	StoryNodeList []StoryNode
	Background    StoryBackground
	Status        RunStatus
	AttrNodeList  []AttrNode
}

func RoomArrLoad() []Room {
	var res []Room
	roomIdArr := RoomIdArrLoad()
	for i := 0; i < len(roomIdArr); i++ {
		res = append(res, RoomLoad(roomIdArr[i]))
	}
	return res
}

func RoomArrSave(roomArr []Room) {
	for i := 0; i < len(roomArr); i++ {
		RoomSave(roomArr[i])
	}
}

func RoomIdArrLoad() []int {
	var res []int
	roomIdArrAddr := "./file/room/room_id_arr.json"
	f, err := ioutil.ReadFile(roomIdArrAddr)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &res)
	if err != nil {
		panic(err)
	}
	return res
}

func RoomIdArrSave(roomIdArr []int) {
	roomIdArrAddr := "./file/room/room_id_arr.json"
	str, err := json.Marshal(roomIdArr)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(roomIdArrAddr, str, 0644)
}

func RoomLoad(roomId int) Room {
	var res Room
	roomAddr := "./file/room/" + strconv.Itoa(roomId)
	res.RoomId = roomId
	roomStoryAddr := roomAddr + "/story.json"
	res.StoryNodeList = StoryLoad(roomStoryAddr)
	roomStatusAddr := roomAddr + "/status.json"
	res.Status = StatusLoad(roomStatusAddr)
	roomBackgroundAddr := roomAddr + "/story_background.json"
	res.Background = StoryBackgroundLoad(roomBackgroundAddr)
	roomAttrAddr := roomAddr + "/attr.json"
	res.AttrNodeList = AttrLoad(roomAttrAddr)
	return res
}

func RoomSave(room Room) {
	roomAddr := "./file/room/" + strconv.Itoa(room.RoomId)
	dirCreate(roomAddr)
	storyAddr := roomAddr + "/story.json"
	StorySave(room.StoryNodeList, storyAddr)
	backgroundAddr := roomAddr + "/story_background.json"
	StoryBackgroundSave(room.Background, backgroundAddr)
	statusAddr := roomAddr + "/status.json"
	StatusSave(room.Status, statusAddr)
	attrAddr := roomAddr + "/attr.json"
	AttrSave(attrAddr, room.AttrNodeList)
}
