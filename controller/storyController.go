package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
	底层思想是操作一个故事表，根据输入输出节点决定故事走向
*/

// 输入输出结构体
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

type StoryNodeMap map[int]*StoryNode

// 开辟空间后暂时不会被回收
var NodeMap StoryNodeMap
var NodeArr []StoryNode

func StoryCreate() {
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
func StoryNodeAdd(val string, input []StorySeleter, output []StorySeleter) bool {
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
func StoryNodeLink(val string, linkInput StorySeleter, linkOutput StorySeleter) bool {
	// 生成新节点
	ok := StoryNodeAdd(val, []StorySeleter{linkInput}, []StorySeleter{linkOutput})
	if !ok {
		return false
	}
	// 断开旧链接
	for i := 0; i < len(NodeMap[linkInput.Id].Output); i++ {
		if linkOutput == NodeMap[linkInput.Id].Output[i] {
			// 删除连接输入节点的原输出节点
			NodeMap[linkInput.Id].Output = append(NodeMap[linkInput.Id].Output[:i], NodeMap[linkInput.Id].Output[i+1:]...)
		}
	}
	for i := 0; i < len(NodeMap[linkOutput.Id].Input); i++ {
		if linkInput == NodeMap[linkOutput.Id].Input[i] {
			// 删除连接输出节点的原输入节点
			NodeMap[linkOutput.Id].Input = append(NodeMap[linkOutput.Id].Input[:i], NodeMap[linkOutput.Id].Input[i+1:]...)
		}
	}
	storySave(NodeArr)
	return true
}

// 故事节点修改
func StoryNodeEdit(nodeId int, val string, input []StorySeleter, output []StorySeleter) bool {
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
			index = searchSelecterId(NodeMap[NodeMap[nodeId].Input[i].Id].Output, nodeId)
			// 从输入节点中删除该节点
			if index != -1 {
				NodeMap[NodeMap[nodeId].Input[i].Id].Output = deleteSelecter(NodeMap[NodeMap[nodeId].Input[i].Id].Output, index)
			}
		}
		// 遍历当前节点的所有输出节点
		for i := 0; i < len(NodeMap[nodeId].Output); i++ {
			index = searchSelecterId(NodeMap[NodeMap[nodeId].Output[i].Id].Input, nodeId)
			// 从输出节点中删除该节点
			if index != -1 {
				NodeMap[NodeMap[nodeId].Output[i].Id].Input = deleteSelecter(NodeMap[NodeMap[nodeId].Output[i].Id].Input, index)
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
	if !storyCheckSelecter(node.Input, NodeMap) || !storyCheckSelecter(node.Output, NodeMap) {
		return false
	}
	for _, v := range node.Input {
		if searchSelecterId(NodeMap[v.Id].Output, node.Id) == -1 {
			NodeMap[v.Id].Output = append(NodeMap[v.Id].Output, StorySeleter{Id: node.Id})
		}
	}
	for _, v := range node.Output {
		if searchSelecterId(NodeMap[v.Id].Input, node.Id) == -1 {
			NodeMap[v.Id].Input = append(NodeMap[v.Id].Input, StorySeleter{Id: node.Id})
		}
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

func storySave(nodeArr []StoryNode) {
	str, err := json.Marshal(nodeArr)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("file/story_example.json", str, 0644)
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

func searchNodeId(nodeArr []StoryNode, id int) int {
	for i := 0; i < len(nodeArr); i++ {
		if id == nodeArr[i].Id {
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

func updateNodeMap() {
	for i := 0; i < len(NodeArr); i++ {
		NodeMap[NodeArr[i].Id] = &NodeArr[i]
	}
}
