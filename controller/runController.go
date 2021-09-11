package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	底层思想是维护一个当前跑团状态表，用以记录当前跑团的路径和经过
*/

type RunStatus struct {
	NowStoryNode    int
	RecordStoryNode []int
}

var Status RunStatus

// 跑团状态新建
func RunStatusCreate() {
	status_new := RunStatus{
		NowStoryNode:    0,
		RecordStoryNode: []int{0},
	}
	Status = status_new
	runStatusSave(status_new)
}

func FileCheck(path string, function func()) {
	_, err := os.Stat(path)
	if err != nil {
		function()
	}
}

// 跑团状态初始化
func RunInit() {
	f, err := ioutil.ReadFile("file/status_example.json")
	if err != nil {
		fmt.Println("read fail", err)
	}
	err = json.Unmarshal(f, &Status)
	if err != nil {
		fmt.Println("json decode fail", err)
	}
}

// 跑团状态展示
func RunStatusList() RunStatus {
	return Status
}

// 跑团当前节点查询
func RunNowNodeGet() StoryNode {
	return StoryNodeGet(Status.NowStoryNode)
}

// 跑团已经过节点查询
func RunNowRecordList() []StoryNode {
	var res []StoryNode
	for i := 0; i < (len(Status.RecordStoryNode) - 1); i++ {
		res = append(res, *StoryNodeMap[Status.RecordStoryNode[i]])
	}
	return res
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
