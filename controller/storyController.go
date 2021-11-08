package controller

import (
	"fmt"
)

/*
	底层思想是操作一个故事表，根据输入输出节点决定故事走向
*/
// TODO 整合故事与背景结构体
// 以storyNodeArr/storyNodeMap为操作实体
// 新的保存方式

type StoryTotalInfo struct {
	StoryTotal []StoryNode
	Background StoryBackground
}

// 输入输出结构体
// 可以添加ext扩展，使run控制器根据条件显示故事选项
type StorySeleter struct {
	Id  int
	Val string
}

type StoryNode struct {
	Id     int
	Val    string
	Input  []StorySeleter // []int
	Output []StorySeleter // []int
}

type StoryBackground struct {
	Background string
}

// 开辟空间后暂时不会被回收
// var StoryNodeMap map[int]*StoryNode
// var StoryNodeArr []StoryNode
// var StoryBackgroundNode StoryBackground

func Test() {
	// RoomCreate()
	fmt.Println("test")
}

func StoryCreate() []StoryNode {
	storyArr := make([]StoryNode, 2)
	storyArr[0] = StoryNode{
		Id:     0,
		Input:  []StorySeleter{},
		Output: []StorySeleter{{Id: 1, Val: "end"}},
		Val:    "start",
	}
	storyArr[1] = StoryNode{
		Id:     1,
		Input:  []StorySeleter{{Id: 0, Val: "start"}},
		Output: []StorySeleter{},
		Val:    "end",
	}
	return storyArr
}

// 故事整体展示
func StoryList(storyNodeArr []StoryNode) []StoryNode {
	return storyNodeArr
}

// 故事节点获取
func StoryNodeGet(storyNodeMap map[int]*StoryNode, id int) StoryNode {
	if data, ok := storyNodeMap[id]; ok {
		return *data
	} else {
		return StoryNode{}
	}
}

func StoryNodeCreate(storyNodeArr []StoryNode, val string) StoryNode {
	var node = StoryNode{
		Id:     0,
		Input:  nil,
		Output: nil,
		Val:    val,
	}
	for _, value := range storyNodeArr {
		if value.Id > node.Id {
			node.Id = value.Id
		}
	}
	node.Id += 1

	return node
}

// 新增故事节点 TODO 输入输出组节点非必须
func StoryNodeAdd(storyNodeArr []StoryNode, storyNodeMap map[int]*StoryNode, val string, input []StorySeleter, output []StorySeleter) bool {
	node := StoryNodeCreate(storyNodeArr, val)
	// append后StoryNodeMap的地址和append的地址不同, 需要更新StoryNodeMap
	storyNodeArr = append(storyNodeArr, node)
	updateNodeMap(storyNodeArr, storyNodeMap)
	// 自己的输出组
	for i := 0; i < len(output); i++ {
		StorySelecterAdd(storyNodeMap, node.Id, output[i].Id, output[i].Val)
	}
	// 对方的输出组
	for i := 0; i < len(input); i++ {
		StorySelecterAdd(storyNodeMap, input[i].Id, node.Id, input[i].Val)
	}

	return true
}

// 故事节点修改
// TODO 选择组非必须
// 更新投票节点
// 如果节点已跑，输出节点不可改
func StoryNodeEdit(storyNodeMap map[int]*StoryNode, nodeId int, val string, input []StorySeleter, output []StorySeleter) bool {
	if data, ok := storyNodeMap[nodeId]; !ok {
		return false
	} else {
		// 先清掉链接
		// 清除自己的输出组
		for i := 0; i < len(data.Output); i++ {
			StorySelecterDelete(storyNodeMap, nodeId, data.Output[i].Id)
		}
		// 清除别人的输出组
		for i := 0; i < len(data.Input); i++ {
			StorySelecterDelete(storyNodeMap, data.Input[i].Id, nodeId)
		}
	}
	// 设置自己的输出组
	for i := 0; i < len(output); i++ {
		StorySelecterAdd(storyNodeMap, nodeId, output[i].Id, output[i].Val)
	}
	// 对方的输出组
	for i := 0; i < len(input); i++ {
		StorySelecterAdd(storyNodeMap, input[i].Id, nodeId, input[i].Val)
	}
	storyNodeMap[nodeId].Val = val
	return true
}

// 故事节点删除
func StoryNodeDelete(storyNodeMap map[int]*StoryNode, storyNodeArr []StoryNode, nodeId int) {
	if nodeId == 1 || nodeId == 0 {
		return
	}
	data, ok := storyNodeMap[nodeId]
	if ok {
		var index int
		// 清除自己的输出组
		for i := 0; i < len(data.Output); i++ {
			StorySelecterDelete(storyNodeMap, nodeId, data.Output[i].Id)
		}
		// 清除别人的输出组
		for i := 0; i < len(data.Input); i++ {
			StorySelecterDelete(storyNodeMap, data.Input[i].Id, nodeId)
		}
		delete(storyNodeMap, nodeId)
		index = searchStoryId(storyNodeArr, nodeId)
		storyNodeArr = deleteNodeSlice(storyNodeArr, index)
	}
}

// 故事节点链接添加
func StorySelecterAdd(storyNodeMap map[int]*StoryNode, nodeId int, linkId int, val string) bool {
	// 新增输出节点
	if node, ok := storyNodeMap[nodeId]; ok {
		index := searchSelecterId(node.Output, linkId)
		if index == -1 {
			node.Output = append(node.Output, StorySeleter{Id: linkId, Val: val})
		} else {
			node.Output[index].Val = val
		}
	} else {
		return false
	}
	// 被输入的节点新增输入节点
	if node, ok := storyNodeMap[linkId]; ok {
		index := searchSelecterId(node.Input, nodeId)
		if index == -1 {
			node.Input = append(node.Input, StorySeleter{Id: nodeId, Val: val})
		} else {
			node.Input[index].Val = val
		}
	} else {
		return false
	}

	return true
}

// 故事节点链接删除
// 查找自己的输出组删除对方输入组
// 查找自己的输入组让对方删除自己
func StorySelecterDelete(storyNodeMap map[int]*StoryNode, nodeId int, linkId int) bool {
	// 删除输出组链接
	if node, ok := storyNodeMap[nodeId]; ok {
		index := searchSelecterId(node.Output, linkId)
		if index != -1 {
			node.Output = deleteSelecter(node.Output, index)
		}
	} else {
		return false
	}
	// 删除输入组链接
	if node, ok := storyNodeMap[linkId]; ok {
		index := searchSelecterId(node.Input, nodeId)
		if index != -1 {
			node.Input = deleteSelecter(node.Input, index)
		}
	} else {
		return false
	}

	return true
}

// 检查输入输出节点是否存在
func storyCheckSelecter(arr []StorySeleter, arrMap map[int]*StoryNode) bool {
	for _, v := range arr {
		_, ok := arrMap[v.Id]
		if !ok {
			return false
		}
	}
	return true
}

func searchSelecterId(idArr []StorySeleter, id int) int {
	for i := 0; i < len(idArr); i++ {
		if id == idArr[i].Id {
			return i
		}
	}
	return -1
}

func searchId(idArr []int, id int) int {
	for i := 0; i < len(idArr); i++ {
		if id == idArr[i] {
			return i
		}
	}
	return -1
}

func searchStoryId(arr []StoryNode, id int) int {
	for i := 0; i < len(arr); i++ {
		if id == arr[i].Id {
			return i
		}
	}
	return -1
}

func deleteSelecter(arr []StorySeleter, index int) []StorySeleter {
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func deleteNodeSlice(arr []StoryNode, index int) []StoryNode {
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func updateNodeMap(storyNodeArr []StoryNode, storyNodeMap map[int]*StoryNode) {
	for i := 0; i < len(storyNodeArr); i++ {
		storyNodeMap[storyNodeArr[i].Id] = &storyNodeArr[i]
	}
}
