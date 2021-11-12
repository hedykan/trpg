package controller

import (
	model "github.com/trpg/model"
)

// room 记录故事节点和状态节点
type Room struct {
	RoomId        int
	StoryNodeList []StoryNode
	StoryNodeMap  map[int]*StoryNode
	Background    StoryBackground
	Status        RunStatus
}

var RoomArr []Room
var RoomMap map[int]*Room
var RoomIdArr []int

// 根据持久化数据初始化房间
func RoomInit() {
	RoomIdArr = model.RoomIdArrLoad()
	RoomArr = RoomArrTransfer(model.RoomArrLoad())
	RoomMap = make(map[int]*Room)
	updateRoomMap(RoomArr, RoomMap)
}

// load/save函数

// 创建房间
func RoomCreate() {
	roomId := roomIdCreate(RoomArr) + 1
	RoomArr = append(RoomArr, Room{
		RoomId:        roomId,
		StoryNodeList: StoryCreate(),
		StoryNodeMap:  map[int]*StoryNode{},
		Background:    StoryBackground{Background: ""},
	})

	RoomIdArr = append(RoomIdArr, roomId)
	updateRoomMap(RoomArr, RoomMap)
	updateNodeMap(RoomMap[roomId].StoryNodeList, RoomMap[roomId].StoryNodeMap)
	RoomMap[roomId].Status = *RunStatusCreate(RoomMap[roomId].StoryNodeMap)

	go model.RoomArrSave(RoomArrTransferModel(RoomArr))
	go model.RoomIdArrSave(RoomIdArr)
}

// 查询房间
type RoomRes struct {
	RoomId     int
	Background string
}

func RoomList() []RoomRes {
	var res []RoomRes
	for i := 0; i < len(RoomArr); i++ {
		res = append(res, RoomRes{
			RoomId:     RoomArr[i].RoomId,
			Background: RoomArr[i].Background.Background,
		})
	}
	return res
}

// Story
// 查询故事
func RoomStoryList(roomId int) []StoryNode {
	return RoomMap[roomId].StoryNodeList
}

// 查询故事节点
func RoomStoryNodeGet(roomId int, nodeId int) StoryNode {
	return StoryNodeGet(RoomMap[roomId].StoryNodeMap, nodeId)
}

// 查询故事背景
func RoomRunBackgroundGet(roomId int) StoryBackground {
	return RoomMap[roomId].Background
}

// 新增房间故事节点
func RoomStoryNodeAdd(roomId int, val string, input []StorySeleter, output []StorySeleter) bool {
	ok := StoryNodeAdd(
		RoomMap[roomId].StoryNodeList,
		RoomMap[roomId].StoryNodeMap,
		val,
		input,
		output,
	)
	go model.RoomSave(RoomTransferModel(*RoomMap[roomId]))
	return ok
}

// 编辑房间故事节点
func RoomStoryNodeEdit(roomId int, nodeId int, val string, input []StorySeleter, output []StorySeleter) bool {
	ok := StoryNodeEdit(
		RoomMap[roomId].StoryNodeMap,
		nodeId,
		val,
		input,
		output,
	)
	go model.RoomSave(RoomTransferModel(*RoomMap[roomId]))
	return ok
}

// 删除房间故事节点
func RoomStoryNodeDelete(roomId int, nodeId int) {
	StoryNodeDelete(
		RoomMap[roomId].StoryNodeMap,
		RoomMap[roomId].StoryNodeList,
		nodeId,
	)
	go model.RoomSave(RoomTransferModel(*RoomMap[roomId]))
}

// 房间故事节点连接
func RoomStorySelecterAdd(roomId int, nodeId int, linkId int, val string) bool {
	ok := StorySelecterAdd(
		RoomMap[roomId].StoryNodeMap,
		nodeId,
		linkId,
		val,
	)
	go model.RoomSave(RoomTransferModel(*RoomMap[roomId]))
	return ok
}

func RoomStorySelecterDelete(roomId int, nodeId int, linkId int) bool {
	ok := StorySelecterDelete(
		RoomMap[roomId].StoryNodeMap,
		nodeId,
		linkId,
	)
	go model.RoomSave(RoomTransferModel(*RoomMap[roomId]))
	return ok
}

// Run
func RoomRunStatusList(roomId int) RunStatus {
	return RunStatusList(
		RoomMap[roomId].Status,
	)
}

// 当前节点获取
func RoomRunNowNodeGet(roomId int) StoryNode {
	return RunNowNodeGet(
		RoomMap[roomId].StoryNodeMap,
		&RoomMap[roomId].Status,
	)
}

// 当前投票获取
func RoomRunNowVoteGet(roomId int) VoteRes {
	return RunNowVoteGet(
		RoomMap[roomId].StoryNodeMap,
		&RoomMap[roomId].Status,
	)
}

// 获取已记录列表
func RoomRunNowRecordList(roomId int) []StoryNode {
	return RunNowRecordList(
		RoomMap[roomId].StoryNodeMap,
		&RoomMap[roomId].Status,
	)
}

// 新投票添加
func RoomRunVoteAdd(roomId int, selecterId int, token string) bool {
	ok := RunVoteAdd(
		RoomMap[roomId].StoryNodeMap,
		&RoomMap[roomId].Status,
		selecterId,
		token,
	)
	go model.RoomSave(RoomTransferModel(*RoomMap[roomId]))
	return ok
}

// 跑团步骤执行
func RoomRunStep(roomId int, nodeId int) {
	RunStep(
		RoomMap[roomId].StoryNodeMap,
		&RoomMap[roomId].Status,
		nodeId,
	)
	model.RoomSave(RoomTransferModel(*RoomMap[roomId]))
}

// 跑团步骤回退
func RoomRunReturn(roomId int, nodeId int) {
	RunReturn(
		RoomMap[roomId].StoryNodeMap,
		&RoomMap[roomId].Status,
		nodeId,
	)
	go model.RoomSave(RoomTransferModel(*RoomMap[roomId]))
}

func roomIdCreate(roomArr []Room) int {
	max := 0
	for i := 0; i < len(roomArr); i++ {
		if roomArr[i].RoomId > max {
			max = roomArr[i].RoomId
		}
	}
	return max
}

func updateRoomMap(roomArr []Room, roomMap map[int]*Room) {
	for i := 0; i < len(roomArr); i++ {
		roomMap[roomArr[i].RoomId] = &roomArr[i]
	}
}
