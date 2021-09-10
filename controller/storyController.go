package controller

import (
	"encoding/json"
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

// 开辟空间后暂时不会被回收
var StoryNodeMap map[int]*StoryNode
var StoryNodeArr []StoryNode

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
	StoryNodeArr = storyArr
	updateNodeMap()
	storySave(storyArr)
}

// 故事整体加载
func StoryLoad() []StoryNode {
	StoryNodeMap = make(map[int]*StoryNode)
	f, err := ioutil.ReadFile("file/story_example.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &StoryNodeArr)
	if err != nil {
		panic(err)
	}
	updateNodeMap()
	return StoryNodeArr
}

// 故事整体展示
func StoryList() []StoryNode {
	return StoryNodeArr
}

// 故事节点获取
func StoryNodeGet(id int) StoryNode {
	return *StoryNodeMap[id]
}

func StoryNodeCreate(val string, input []StorySeleter, output []StorySeleter) StoryNode {
	var node = StoryNode{
		Id:     0,
		Input:  input,
		Output: output,
		Val:    val,
	}
	for _, value := range StoryNodeArr {
		if value.Id > node.Id {
			node.Id = value.Id
		}
	}
	node.Id += 1

	return node
}

// 新增故事节点 TODO 给输入输出节点新增节点
func StoryNodeAdd(val string, input []StorySeleter, output []StorySeleter) bool {
	node := StoryNodeCreate(val, input, output)

	ok := storyInOutSet(node)
	if !ok {
		return false
	}
	// append后StoryNodeMap的地址和append的地址不同, 需要更新StoryNodeMap
	StoryNodeArr = append(StoryNodeArr, node)
	updateNodeMap()
	// 存储新节点
	storySave(StoryNodeArr)

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
	for i := 0; i < len(StoryNodeMap[linkInput.Id].Output); i++ {
		if linkOutput == StoryNodeMap[linkInput.Id].Output[i] {
			// 删除连接输入节点的原输出节点
			StoryNodeMap[linkInput.Id].Output = append(StoryNodeMap[linkInput.Id].Output[:i], StoryNodeMap[linkInput.Id].Output[i+1:]...)
		}
	}
	for i := 0; i < len(StoryNodeMap[linkOutput.Id].Input); i++ {
		if linkInput == StoryNodeMap[linkOutput.Id].Input[i] {
			// 删除连接输出节点的原输入节点
			StoryNodeMap[linkOutput.Id].Input = append(StoryNodeMap[linkOutput.Id].Input[:i], StoryNodeMap[linkOutput.Id].Input[i+1:]...)
		}
	}
	storySave(StoryNodeArr)
	return true
}

// 故事节点修改
func StoryNodeEdit(nodeId int, val string, input []StorySeleter, output []StorySeleter) bool {
	if _, ok := StoryNodeMap[nodeId]; !ok {
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
	StoryNodeMap[nodeId].Val = val
	StoryNodeMap[nodeId].Input = input
	StoryNodeMap[nodeId].Output = output
	return true
}

// 故事节点删除
func StoryNodeDelete(nodeId int) {
	_, ok := StoryNodeMap[nodeId]
	if ok {
		var index int
		// 遍历当前节点的所有输入节点
		for i := 0; i < len(StoryNodeMap[nodeId].Input); i++ {
			index = searchSelecterId(StoryNodeMap[StoryNodeMap[nodeId].Input[i].Id].Output, nodeId)
			// 从输入节点中删除该节点
			if index != -1 {
				StoryNodeMap[StoryNodeMap[nodeId].Input[i].Id].Output = deleteSelecter(StoryNodeMap[StoryNodeMap[nodeId].Input[i].Id].Output, index)
			}
		}
		// 遍历当前节点的所有输出节点
		for i := 0; i < len(StoryNodeMap[nodeId].Output); i++ {
			index = searchSelecterId(StoryNodeMap[StoryNodeMap[nodeId].Output[i].Id].Input, nodeId)
			// 从输出节点中删除该节点
			if index != -1 {
				StoryNodeMap[StoryNodeMap[nodeId].Output[i].Id].Input = deleteSelecter(StoryNodeMap[StoryNodeMap[nodeId].Output[i].Id].Input, index)
			}
		}
		delete(StoryNodeMap, nodeId)
		index = searchNodeId(StoryNodeArr, nodeId)
		StoryNodeArr = deleteNodeSlice(StoryNodeArr, index)
		storySave(StoryNodeArr)
	}
}

// 设置输入输出节点
func storyInOutSet(node StoryNode) bool {
	if !storyCheckSelecter(node.Input, StoryNodeMap) || !storyCheckSelecter(node.Output, StoryNodeMap) {
		return false
	}
	for _, v := range node.Input {
		if searchSelecterId(StoryNodeMap[v.Id].Output, node.Id) == -1 {
			StoryNodeMap[v.Id].Output = append(StoryNodeMap[v.Id].Output, StorySeleter{Id: node.Id})
		}
	}
	for _, v := range node.Output {
		if searchSelecterId(StoryNodeMap[v.Id].Input, node.Id) == -1 {
			StoryNodeMap[v.Id].Input = append(StoryNodeMap[v.Id].Input, StorySeleter{Id: node.Id})
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

func storySave(arr []StoryNode) {
	str, err := json.Marshal(arr)
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

func searchNodeId(arr []StoryNode, id int) int {
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

func updateNodeMap() {
	for i := 0; i < len(StoryNodeArr); i++ {
		StoryNodeMap[StoryNodeArr[i].Id] = &StoryNodeArr[i]
	}
}
