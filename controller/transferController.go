package controller

import model "github.com/trpg/model"

// Story
func StoryNodeArrTransfer(storyNodeArr []model.StoryNode) []StoryNode {
	var res []StoryNode
	for i := 0; i < len(storyNodeArr); i++ {
		res = append(res, StoryNode{
			Id:     storyNodeArr[i].Id,
			Val:    storyNodeArr[i].Val,
			Input:  StorySelecterArrTransfer(storyNodeArr[i].Input),
			Output: StorySelecterArrTransfer(storyNodeArr[i].Output),
		})
	}
	return res
}

func StoryNodeArrTransferModel(storyNodeArr []StoryNode) []model.StoryNode {
	var res []model.StoryNode
	for i := 0; i < len(storyNodeArr); i++ {
		res = append(res, model.StoryNode{
			Id:     storyNodeArr[i].Id,
			Val:    storyNodeArr[i].Val,
			Input:  StorySelecterArrTransferModel(storyNodeArr[i].Input),
			Output: StorySelecterArrTransferModel(storyNodeArr[i].Output),
		})
	}
	return res
}

func StorySelecterArrTransfer(storySelecterArr []model.StorySeleter) []StorySeleter {
	var res []StorySeleter
	for i := 0; i < len(storySelecterArr); i++ {
		res = append(res, StorySeleter{
			Id:  storySelecterArr[i].Id,
			Val: storySelecterArr[i].Val,
		})
	}
	return res
}

func StorySelecterArrTransferModel(storySelecterArr []StorySeleter) []model.StorySeleter {
	var res []model.StorySeleter
	for i := 0; i < len(storySelecterArr); i++ {
		res = append(res, model.StorySeleter{
			Id:  storySelecterArr[i].Id,
			Val: storySelecterArr[i].Val,
		})
	}
	return res
}

func StoryBackgroundTransfer(storyBackground model.StoryBackground) StoryBackground {
	var res StoryBackground
	res = StoryBackground{
		Background: storyBackground.Background,
	}
	return res
}

func StoryBackgroundTransferModel(storyBackground StoryBackground) model.StoryBackground {
	var res model.StoryBackground
	res = model.StoryBackground{
		Background: storyBackground.Background,
	}
	return res
}

// Run
func StatusTransfer(status model.RunStatus) RunStatus {
	var res RunStatus
	res = RunStatus{
		NowStoryNode:    status.NowStoryNode,
		RecordStoryNode: status.RecordStoryNode,
		RecordVote:      VoteTransfer(status.RecordVote),
	}
	return res
}

func StatusTransferModel(status RunStatus) model.RunStatus {
	var res model.RunStatus
	res = model.RunStatus{
		NowStoryNode:    status.NowStoryNode,
		RecordStoryNode: status.RecordStoryNode,
		RecordVote:      VoteTransferModel(status.RecordVote),
	}
	return res
}

func VoteTransfer(voteArr []model.RunVote) []RunVote {
	var res []RunVote
	for i := 0; i < len(voteArr); i++ {
		res = append(res, RunVote{
			NodeId:         voteArr[i].NodeId,
			VoteStatusList: VoteStatusArrTransfer(voteArr[i].VoteStatusList),
			TokenList:      voteArr[i].TokenList,
		})
	}
	return res
}

func VoteTransferModel(voteArr []RunVote) []model.RunVote {
	var res []model.RunVote
	for i := 0; i < len(voteArr); i++ {
		res = append(res, model.RunVote{
			NodeId:         voteArr[i].NodeId,
			VoteStatusList: VoteStatusArrTransferModel(voteArr[i].VoteStatusList),
			TokenList:      voteArr[i].TokenList,
		})
	}
	return res
}

func VoteStatusArrTransfer(voteStatusArr []model.RunVoteStatus) []RunVoteStatus {
	var res []RunVoteStatus
	for i := 0; i < len(voteStatusArr); i++ {
		res = append(res, RunVoteStatus{
			SelecterId: voteStatusArr[i].SelecterId,
			Num:        voteStatusArr[i].Num,
		})
	}
	return res
}

func VoteStatusArrTransferModel(voteStatusArr []RunVoteStatus) []model.RunVoteStatus {
	var res []model.RunVoteStatus
	for i := 0; i < len(voteStatusArr); i++ {
		res = append(res, model.RunVoteStatus{
			SelecterId: voteStatusArr[i].SelecterId,
			Num:        voteStatusArr[i].Num,
		})
	}
	return res
}

// attr
func AttrArrTransfer(attrArr []model.AttrNode) []AttrNode {
	var res []AttrNode
	for i := 0; i < len(attrArr); i++ {
		res = append(res, AttrNode{
			Id:  attrArr[i].Id,
			Val: attrArr[i].Val,
			Num: attrArr[i].Num,
		})
	}
	return res
}

func AttrArrTransferModel(attrArr []AttrNode) []model.AttrNode {
	var res []model.AttrNode
	for i := 0; i < len(attrArr); i++ {
		res = append(res, model.AttrNode{
			Id:  attrArr[i].Id,
			Val: attrArr[i].Val,
			Num: attrArr[i].Num,
		})
	}
	return res
}

// Room
// 接收请求回来的数据
func RoomArrTransfer(roomArr []model.Room) []Room {
	var res []Room
	for i := 0; i < len(roomArr); i++ {
		res = append(res, RoomTransfer(roomArr[i]))
	}
	return res
}

func RoomArrTransferModel(roomArr []Room) []model.Room {
	var res []model.Room
	for i := 0; i < len(roomArr); i++ {
		res = append(res, RoomTransferModel(roomArr[i]))
	}
	return res
}

func RoomTransfer(room model.Room) Room {
	var res Room
	res = Room{
		RoomId:     room.RoomId,
		Story:      StoryTable{StoryList: StoryNodeArrTransfer(room.StoryNodeList), StoryMap: make(map[int]*StoryNode)},
		Background: StoryBackgroundTransfer(room.Background),
		Status:     StatusTransfer(room.Status),
		Attribute:  AttrTable{AttrList: AttrArrTransfer(room.AttrNodeList), AttrMap: make(map[int]*AttrNode)},
	}
	res.Story.updateMap()
	res.Attribute.updateMap()
	return res
}

func RoomTransferModel(room Room) model.Room {
	var res model.Room
	res = model.Room{
		RoomId:        room.RoomId,
		StoryNodeList: StoryNodeArrTransferModel(room.Story.StoryList),
		Background:    model.StoryBackground(room.Background),
		Status:        StatusTransferModel(room.Status),
		AttrNodeList:  AttrArrTransferModel(room.Attribute.AttrList),
	}
	return res
}
