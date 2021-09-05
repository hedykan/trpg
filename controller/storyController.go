package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
	底层思想是操作一个故事表，根据输入输出节点决定故事走向
*/

type StoryNode struct {
	Id     int
	Val    string
	Input  []int
	Output []int
}

type StoryNodeMap map[int]*StoryNode

// 开辟空间后暂时不会被回收
var NodeMap StoryNodeMap
var NodeArr []StoryNode

func StoryCreate() {
	storyArr := make([]StoryNode, 2)
	storyArr[0] = StoryNode{
		Id:     0,
		Input:  []int{},
		Output: []int{1},
		Val:    "start",
	}
	storyArr[1] = StoryNode{
		Id:     1,
		Input:  []int{0},
		Output: []int{},
		Val:    "end",
	}
	NodeArr = storyArr
	for i := 0; i < len(NodeArr); i++ {
		NodeMap[NodeArr[i].Id] = &NodeArr[i]
	}
	storySave(storyArr)
}

// 故事整体初始化
func StoryInit() []StoryNode {
	NodeMap = make(map[int]*StoryNode)
	f, err := ioutil.ReadFile("file/story_example.json")
	if err != nil {
		fmt.Println("read fail", err)
	}
	err = json.Unmarshal(f, &NodeArr)
	if err != nil {
		fmt.Println("json decode fail", err)
	}
	updateNodeMap()
	return NodeArr
}

// 故事整体展示
func StoryList() []StoryNode {
	return NodeArr
}

// 故事节点获取
func StoryNodeGet(id int) StoryNode {
	return *NodeMap[id]
}

// 新增故事节点 TODO 给输入输出节点新增节点
func StoryNodeAdd(val string, input []int, output []int) bool {
	var Node StoryNode
	Node = StoryNode{
		Id:     0,
		Input:  input,
		Output: output,
		Val:    val,
	}
	for _, value := range NodeArr {
		if value.Id > Node.Id {
			Node.Id = value.Id
		}
	}
	Node.Id += 1

	ok := storyInOutSet(Node)
	if !ok {
		return false
	}
	// append后NodeMap的地址和append的地址不同, 需要更新NodeMap
	NodeArr = append(NodeArr, Node)
	updateNodeMap()
	// 存储新节点
	storySave(NodeArr)

	return true
}

// 插入链接故事节点
func StoryNodeLink(val string, linkInput int, linkOutput int) bool {
	// 生成新节点
	ok := StoryNodeAdd(val, []int{linkInput}, []int{linkOutput})
	if !ok {
		return false
	}
	// 断开旧链接
	for i := 0; i < len(NodeMap[linkInput].Output); i++ {
		if linkOutput == NodeMap[linkInput].Output[i] {
			// 删除连接输入节点的原输出节点
			NodeMap[linkInput].Output = append(NodeMap[linkInput].Output[:i], NodeMap[linkInput].Output[i+1:]...)
		}
	}
	for i := 0; i < len(NodeMap[linkOutput].Input); i++ {
		if linkInput == NodeMap[linkOutput].Input[i] {
			// 删除连接输出节点的原输入节点
			NodeMap[linkOutput].Input = append(NodeMap[linkOutput].Input[:i], NodeMap[linkOutput].Input[i+1:]...)
		}
	}
	storySave(NodeArr)
	return true
}

// 故事节点修改
func StoryNodeEdit(nodeId int, val string, input []int, output []int) bool {
	if _, ok := NodeMap[nodeId]; !ok {
		return false
	}
	var node StoryNode
	node = StoryNode{
		Id:     nodeId,
		Input:  input,
		Output: output,
		Val:    val,
	}
	ok := storyInOutSet(node)
	if !ok {
		return false
	}
	NodeMap[nodeId].Val = val
	NodeMap[nodeId].Input = input
	NodeMap[nodeId].Output = output
	return true
}

// 故事节点删除
func StoryNodeDelete(nodeId int) {
	_, ok := NodeMap[nodeId]
	if ok {
		var index int
		// 遍历当前节点的所有输入节点
		for i := 0; i < len(NodeMap[nodeId].Input); i++ {
			index = searchId(NodeMap[NodeMap[nodeId].Input[i]].Output, nodeId)
			// 从输入节点中删除该节点
			if index != -1 {
				NodeMap[NodeMap[nodeId].Input[i]].Output = deleteIntSlice(NodeMap[NodeMap[nodeId].Input[i]].Output, index)
			}
		}
		// 遍历当前节点的所有输出节点
		for i := 0; i < len(NodeMap[nodeId].Output); i++ {
			index = searchId(NodeMap[NodeMap[nodeId].Output[i]].Input, nodeId)
			// 从输出节点中删除该节点
			if index != -1 {
				NodeMap[NodeMap[nodeId].Output[i]].Input = deleteIntSlice(NodeMap[NodeMap[nodeId].Output[i]].Input, index)
			}
		}
		delete(NodeMap, nodeId)
		index = searchNodeId(NodeArr, nodeId)
		NodeArr = deleteNodeSlice(NodeArr, index)
		storySave(NodeArr)
	}
}

// 设置输入输出节点
func storyInOutSet(node StoryNode) bool {
	if !storyCheckInOut(node.Input, NodeMap) || !storyCheckInOut(node.Output, NodeMap) {
		return false
	}
	for _, v := range node.Input {
		if searchId(NodeMap[v].Output, node.Id) == -1 {
			NodeMap[v].Output = append(NodeMap[v].Output, node.Id)
		}
	}
	for _, v := range node.Output {
		if searchId(NodeMap[v].Input, node.Id) == -1 {
			NodeMap[v].Input = append(NodeMap[v].Input, node.Id)
		}
	}
	return true
}

// 检查输入输出节点是否存在
func storyCheckInOut(arr []int, arrMap map[int]*StoryNode) bool {
	for _, v := range arr {
		_, ok := arrMap[v]
		if !ok {
			return false
		}
	}
	return true
}

func storySave(nodeArr []StoryNode) {
	str, err := json.Marshal(nodeArr)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("file/story_example.json", str, 0644)
}

func searchId(idArr []int, id int) int {
	for i := 0; i < len(idArr); i++ {
		if id == idArr[i] {
			return i
		}
	}
	return -1
}

func searchNodeId(nodeArr []StoryNode, id int) int {
	for i := 0; i < len(nodeArr); i++ {
		if id == nodeArr[i].Id {
			return i
		}
	}
	return -1
}

func deleteIntSlice(arr []int, index int) []int {
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func deleteNodeSlice(arr []StoryNode, index int) []StoryNode {
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func updateNodeMap() {
	for i := 0; i < len(NodeArr); i++ {
		NodeMap[NodeArr[i].Id] = &NodeArr[i]
	}
}
