package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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

func RunNowNodeGet() StoryNode {
	return StoryNodeGet(Status.NowStoryNode)
}

// 步骤执行
func RunStep(nodeId int) {
	// 确定有当地故事节点可以进入目标节点
	ok := searchId(NodeMap[Status.NowStoryNode].Output, nodeId)
	if ok == -1 {
		return
	}
	// 设置状态
	Status.NowStoryNode = nodeId
	Status.RecordStoryNode = append(Status.RecordStoryNode, nodeId)
	runStatusSave(Status)
}

// 保存跑团状态
func runStatusSave(status RunStatus) {
	str, err := json.Marshal(status)
	if err != nil {
		fmt.Println("transfer err", err)
	}
	ioutil.WriteFile("file/status_example.json", str, 0644)
}
