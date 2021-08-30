package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 内容名称要大写才能导出
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
	for i := 0; i < len(NodeArr); i++ {
		NodeMap[NodeArr[i].Id] = &NodeArr[i]
	}
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

// 新增故事节点
func StoryNodeAdd(val string, input []int, output []int) int {
	var node StoryNode
	node.Id = 0
	node.Val = val
	node.Input = input
	node.Output = output
	for _, value := range NodeArr {
		if value.Id > node.Id {
			node.Id = value.Id
		}
	}
	node.Id += 1

	// 存储新节点
	NodeMap[node.Id] = &node
	NodeArr = append(NodeArr, node)
	storySave(NodeArr)

	return node.Id
}

// 插入链接故事节点
func StoryNodeLink(val string, linkInput int, linkOutput int) {
	// 生成新节点
	storyId := StoryNodeAdd(val, []int{linkInput}, []int{linkOutput})
	// 断开旧链接
	NodeMap[linkInput].Output = append(NodeMap[linkInput].Output, storyId)
	NodeMap[linkOutput].Input = append(NodeMap[linkOutput].Input, storyId)
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
}

func storySave(nodeArr []StoryNode) {
	str, err := json.Marshal(nodeArr)
	if err != nil {
		fmt.Println("transfer err", err)
	}
	ioutil.WriteFile("file/story_example.json", str, 0644)
}

func searchId(idArr []int, id int) int {
	var i int
	for i = 0; i < len(idArr); i++ {
		if id == idArr[i] {
			return i
		}
	}
	return -1
}
