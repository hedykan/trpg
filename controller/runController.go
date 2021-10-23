package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	底层思想是维护一个当前跑团状态表，用以记录当前跑团的路径和经过
	记录每一个节点的投票状态
*/
type RunVoteStatus struct {
	SelecterId int
	Num        int
	TokenList  []string
}

type RunVote struct {
	NodeId         int
	VoteStatusList []RunVoteStatus
}

type RunStatus struct {
	NowStoryNode    int
	RecordStoryNode []int
	RecordVote      []RunVote
}

var Status RunStatus

// 跑团状态新建
func RunStatusCreate() {
	RunVoteNew, err := RunVoteCreate(0)
	if err != nil {
		panic(err)
	}
	Status = RunStatus{
		NowStoryNode:    0,
		RecordStoryNode: []int{0},
		RecordVote:      []RunVote{RunVoteNew},
	}
	runStatusSave(Status)
}

// 跑团加载
func RunLoad() {
	f, err := ioutil.ReadFile("./file/status_example.json")
	if err != nil {
		fmt.Println("read fail", err)
	}
	err = json.Unmarshal(f, &Status)
	if err != nil {
		fmt.Println("json decode fail", err)
	}
}

// 跑团状态初始化
func RunInit() {
	path, _ := os.Getwd()
	path = path + "/file/status_example.json"
	FileCheck(path, RunStatusCreate)
	RunLoad()
}

// 跑团故事背景获取
func RunStoryBackgroundGet() StoryBackground {
	return StoryBackgroundNode
}

// 跑团状态展示
func RunStatusList() RunStatus {
	return Status
}

// 跑团当前节点查询
func RunNowNodeGet() StoryNode {
	return StoryNodeGet(Status.NowStoryNode)
}

func RunNowVoteGet() struct {
	NodeId     int
	VoteStatus []struct {
		SelecterId int
		Val        string
		Num        int
	}
} {
	data := Status.RecordVote
	var res struct {
		NodeId     int
		VoteStatus []struct {
			SelecterId int
			Val        string
			Num        int
		}
	}
	// 查找节点
	// 回复加工后的RunVote
	for i := 0; i < len(data); i++ {
		if data[i].NodeId == Status.NowStoryNode {
			storyNode := RunNowNodeGet()
			find := data[i]
			res.NodeId = Status.NowStoryNode
			res.VoteStatus = make([]struct {
				SelecterId int
				Val        string
				Num        int
			}, len(storyNode.Output))
			for i := 0; i < len(find.VoteStatusList); i++ {
				res.VoteStatus[i].SelecterId = find.VoteStatusList[i].SelecterId
				res.VoteStatus[i].Num = find.VoteStatusList[i].Num
				res.VoteStatus[i].Val = storyNode.Output[i].Val
			}
			break
		}
	}
	return res
}

// 跑团已经过节点查询
func RunNowRecordList() []StoryNode {
	var res []StoryNode
	for i := 0; i < (len(Status.RecordStoryNode) - 1); i++ {
		res = append(res, *StoryNodeMap[Status.RecordStoryNode[i]])
	}
	return res
}

// 节点投票创建
func RunVoteCreate(nodeId int) (RunVote, error) {
	VoteNode := RunVote{}
	if data, ok := StoryNodeMap[nodeId]; ok {
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
func RunVoteAdd(selecterId int, token string) {
	ok := searchSelecterId(StoryNodeMap[Status.NowStoryNode].Output, selecterId)
	if ok == -1 {
		return
	}
	data := &Status.RecordVote[len(Status.RecordVote)-1].VoteStatusList[ok]
	ok = searchToken2List(data.TokenList, token)
	if ok != -1 {
		return
	}
	// 投票+1
	data.Num += 1
	data.TokenList = append(data.TokenList, token)
	runStatusSave(Status)
}

func searchToken2List(tokenArr []string, token string) int {
	for i := 0; i < len(tokenArr); i++ {
		if tokenArr[i] == token {
			return i
		}
	}
	return -1
}

// 步骤执行
func RunStep(nodeId int) {
	// 确定有当地故事节点可以进入目标节点
	ok := searchSelecterId(StoryNodeMap[Status.NowStoryNode].Output, nodeId)
	if ok == -1 {
		return
	}
	// 设置状态
	Status.NowStoryNode = nodeId
	Status.RecordStoryNode = append(Status.RecordStoryNode, nodeId)
	VoteNode, _ := RunVoteCreate(nodeId)
	Status.RecordVote = append(Status.RecordVote, VoteNode)
	runStatusSave(Status)
}

// 步骤回退
func RunReturn(nodeId int) {
	index := searchId(Status.RecordStoryNode, nodeId)
	if index != -1 {
		Status.NowStoryNode = nodeId
		Status.RecordStoryNode = Status.RecordStoryNode[:index+1]
		runStatusSave(Status)
	}
}

// 保存跑团状态
func runStatusSave(status RunStatus) {
	str, err := json.Marshal(status)
	if err != nil {
		fmt.Println("transfer err", err)
	}
	ioutil.WriteFile("file/status_example.json", str, 0644)
}
