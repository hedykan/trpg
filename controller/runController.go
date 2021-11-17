package controller

import (
	"errors"
)

/*
	底层思想是维护一个当前跑团状态表，用以记录当前跑团的路径和经过
	记录每一个节点的投票状态
*/
type RunVoteStatus struct {
	SelecterId int
	Num        int
}

type RunVote struct {
	NodeId         int // 当前节点i
	VoteStatusList []RunVoteStatus
	TokenList      []string
}

type RunStatus struct {
	NowStoryNode    int
	RecordStoryNode []int
	RecordVote      []RunVote
}

// var Status RunStatus

// 跑团状态新建
func RunStatusCreate(storyNodeMap map[int]*StoryNode) *RunStatus {
	runVoteNew, err := RunVoteCreate(storyNodeMap, 0)
	if err != nil {
		panic(err)
	}
	return &RunStatus{
		NowStoryNode:    0,
		RecordStoryNode: []int{0},
		RecordVote:      []RunVote{runVoteNew},
	}
}

// 跑团状态展示
func RunStatusList(status RunStatus) RunStatus {
	return status
}

// 跑团当前节点查询
func RunNowNodeGet(table *StoryTable, status *RunStatus) StoryNode {
	return StoryNodeGet(table, status.NowStoryNode)
}

type VoteRes struct {
	NodeId     int
	VoteStatus []struct {
		Status RunVoteStatus
		Val    string
	}
}

func RunNowVoteGet(table *StoryTable, status *RunStatus) VoteRes {
	data := status.RecordVote
	var res VoteRes
	// 查找节点
	// 回复加工后的RunVote
	for i := 0; i < len(data); i++ {
		if data[i].NodeId == status.NowStoryNode {
			storyNode := RunNowNodeGet(table, status)
			find := data[i]
			res.NodeId = status.NowStoryNode
			res.VoteStatus = make([]struct {
				Status RunVoteStatus
				Val    string
			}, len(storyNode.Output))
			for i := 0; i < len(find.VoteStatusList); i++ {
				res.VoteStatus[i].Status = find.VoteStatusList[i]
				res.VoteStatus[i].Val = storyNode.Output[i].Val
			}
			break
		}
	}
	return res
}

// 跑团经过节点查询
func RunNowRecordList(storyNodeMap map[int]*StoryNode, status *RunStatus) []StoryNode {
	var res []StoryNode
	for i := 0; i < (len(status.RecordStoryNode) - 1); i++ {
		res = append(res, *storyNodeMap[status.RecordStoryNode[i]])
	}
	return res
}

// 节点投票创建
func RunVoteCreate(storyNodeMap map[int]*StoryNode, nodeId int) (RunVote, error) {
	VoteNode := RunVote{}
	if data, ok := storyNodeMap[nodeId]; ok {
		VoteNode.NodeId = data.Id
		// 创建投票列表
		for i := 0; i < len(data.Output); i++ {
			VoteNode.VoteStatusList = append(VoteNode.VoteStatusList, RunVoteStatus{SelecterId: data.Output[i].Id, Num: 0})
		}
		return VoteNode, nil
	} else {
		err := errors.New("node not found")
		return VoteNode, err
	}
}

// 节点id投票
func RunVoteAdd(storyNodeMap map[int]*StoryNode, status *RunStatus, selecterId int, token string) bool {
	index := searchSelecterId(storyNodeMap[status.NowStoryNode].Output, selecterId)
	if index == -1 {
		return false
	}
	data := &status.RecordVote[len(status.RecordVote)-1]
	ok := searchToken2List(data.TokenList, token)
	if ok != -1 {
		return false
	}
	// 投票+1
	data.VoteStatusList[index].Num += 1
	data.TokenList = append(data.TokenList, token)

	return true
}

// 节点id清理
// 根据节点修改后的选择重置
func RunVoteClear(storyNodeMap map[int]*StoryNode, status *RunStatus, nodeId int) {
	index := searchVoteIndex(status.RecordVote, nodeId)
	// 清除目标投票
	if index != -1 {
		status.RecordVote = append(status.RecordVote[:index], status.RecordVote[index+1:]...)
	}
	runVote, err := RunVoteCreate(storyNodeMap, nodeId)
	if err != nil {
		return
	}
	status.RecordVote = append(status.RecordVote, runVote)
}

// 步骤执行
func RunStep(storyNodeMap map[int]*StoryNode, status *RunStatus, nodeId int) {
	// 确定有当地故事节点可以进入目标节点
	ok := searchSelecterId(storyNodeMap[status.NowStoryNode].Output, nodeId)
	if ok == -1 {
		return
	}
	// 设置状态
	status.NowStoryNode = nodeId
	status.RecordStoryNode = append(status.RecordStoryNode, nodeId)
	voteNode, _ := RunVoteCreate(storyNodeMap, nodeId)
	status.RecordVote = append(status.RecordVote, voteNode)
}

// 步骤回退
// 投票也重设
func RunReturn(storyNodeMap map[int]*StoryNode, status *RunStatus, nodeId int) {
	index := searchId(status.RecordStoryNode, nodeId)
	if index != -1 {
		status.NowStoryNode = nodeId
		status.RecordStoryNode = status.RecordStoryNode[:index+1] // 截到目标点
		status.RecordVote = status.RecordVote[:index+1]           // 截到目标点
		RunVoteClear(storyNodeMap, status, nodeId)
	}
}

// 查询token表
func searchToken2List(tokenArr []string, token string) int {
	for i := 0; i < len(tokenArr); i++ {
		if tokenArr[i] == token {
			return i
		}
	}
	return -1
}

func searchVoteIndex(arr []RunVote, nodeId int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i].NodeId == nodeId {
			return i
		}
	}
	return -1
}
