package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	底层思想是操作一个故事表，根据输入输出节点决定故事走向
*/

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
var StoryNodeMap map[int]*StoryNode
var StoryNodeArr []StoryNode
var StoryBackgroundNode StoryBackground

func Test() {
	fmt.Println("test")
}

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
	storySave(storyArr)
}

// 故事整体加载
func StoryLoad() []StoryNode {
	StoryNodeMap = make(map[int]*StoryNode)
	f, err := ioutil.ReadFile("./file/story_example.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &StoryNodeArr)
	if err != nil {
		panic(err)
	}
	return StoryNodeArr
}

func StoryBackgroundLoad() StoryBackground {
	f, err := ioutil.ReadFile("./file/story_background_example.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &StoryBackgroundNode)
	if err != nil {
		panic(err)
	}
	return StoryBackgroundNode
}

func StoryInit() {
	path, _ := os.Getwd()
	path = path + "/file/story_example.json"
	FileCheck(path, StoryCreate)
	StoryLoad()
	updateNodeMap()
}

func StroyBackgroundInit() {
	path, _ := os.Getwd()
	path = path + "/file/story_background_exampld.json"
	FileCheck(path, func() {})
	StoryBackgroundLoad()
}

// 故事整体展示
func StoryList() []StoryNode {
	return StoryNodeArr
}

// 故事节点获取
func StoryNodeGet(id int) StoryNode {
	return *StoryNodeMap[id]
}

func StoryNodeCreate(val string) StoryNode {
	var node = StoryNode{
		Id:     0,
		Input:  nil,
		Output: nil,
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
	node := StoryNodeCreate(val)
	// append后StoryNodeMap的地址和append的地址不同, 需要更新StoryNodeMap
	StoryNodeArr = append(StoryNodeArr, node)
	updateNodeMap()
	// 自己的输出组
	for i := 0; i < len(output); i++ {
		StorySelecterAdd(node.Id, output[i].Id, output[i].Val)
	}
	// 对方的输出组
	for i := 0; i < len(input); i++ {
		StorySelecterAdd(input[i].Id, node.Id, input[i].Val)
	}
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
	if data, ok := StoryNodeMap[nodeId]; !ok {
		return false
	} else {
		// 先清掉链接
		// 清除自己的输出组
		for i := 0; i < len(data.Output); i++ {
			StorySelecterDelete(nodeId, data.Output[i].Id)
		}
		// 清除别人的输出组
		for i := 0; i < len(data.Input); i++ {
			StorySelecterDelete(data.Input[i].Id, nodeId)
		}
	}
	// 设置自己的输出组
	for i := 0; i < len(output); i++ {
		StorySelecterAdd(nodeId, output[i].Id, output[i].Val)
	}
	// 对方的输出组
	for i := 0; i < len(input); i++ {
		StorySelecterAdd(input[i].Id, nodeId, input[i].Val)
	}
	StoryNodeMap[nodeId].Val = val
	return true
}

// 故事节点删除
func StoryNodeDelete(nodeId int) {
	data, ok := StoryNodeMap[nodeId]
	if ok {
		var index int
		// 清除自己的输出组
		for i := 0; i < len(data.Output); i++ {
			StorySelecterDelete(nodeId, data.Output[i].Id)
		}
		// 清除别人的输出组
		for i := 0; i < len(data.Input); i++ {
			StorySelecterDelete(data.Input[i].Id, nodeId)
		}
		delete(StoryNodeMap, nodeId)
		index = searchStoryId(StoryNodeArr, nodeId)
		StoryNodeArr = deleteNodeSlice(StoryNodeArr, index)
		storySave(StoryNodeArr)
	}
}

// 故事节点链接添加
func StorySelecterAdd(nodeId int, linkId int, val string) bool {
	// 新增输出节点
	if node, ok := StoryNodeMap[nodeId]; ok {
		index := searchSelecterId(node.Output, linkId)
		if index == -1 {
			node.Output = append(node.Output, StorySeleter{Id: linkId, Val: val})
		} else {
			node.Output[index].Val = val
		}
	}
	// 被输入的节点新增输入节点
	if node, ok := StoryNodeMap[linkId]; ok {
		index := searchSelecterId(node.Input, nodeId)
		if index == -1 {
			node.Input = append(node.Input, StorySeleter{Id: nodeId, Val: val})
		} else {
			node.Input[index].Val = val
		}
	}
	storySave(StoryNodeArr)
	return true
}

// 故事节点链接删除
// 查找自己的输出组删除对方输入组
// 查找自己的输入组让对方删除自己
func StorySelecterDelete(nodeId int, linkId int) bool {
	// 删除输出组链接
	if node, ok := StoryNodeMap[nodeId]; ok {
		index := searchSelecterId(node.Output, linkId)
		if index != -1 {
			node.Output = deleteSelecter(node.Output, index)
		}
	}
	// 删除输入组链接
	if node, ok := StoryNodeMap[linkId]; ok {
		index := searchSelecterId(node.Input, nodeId)
		if index != -1 {
			node.Input = deleteSelecter(node.Input, index)
		}
	}
	storySave(StoryNodeArr)
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

func updateNodeMap() {
	for i := 0; i < len(StoryNodeArr); i++ {
		StoryNodeMap[StoryNodeArr[i].Id] = &StoryNodeArr[i]
	}
}
