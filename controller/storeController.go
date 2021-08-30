package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 内容名称要大写才能导出
type StoreNode struct {
	Id     int
	Val    string
	Input  []int
	Output []int
}

type storeNodeMap map[int]*StoreNode

// 开辟空间后暂时不会被回收
var NodeMap storeNodeMap
var nodeArr []StoreNode

func StoreInit() []StoreNode {
	NodeMap = make(map[int]*StoreNode)
	f, err := ioutil.ReadFile("file/example.json")
	if err != nil {
		fmt.Println("read fail", err)
	}
	err = json.Unmarshal(f, &nodeArr)
	if err != nil {
		fmt.Println("json decode fail", err)
	}
	for i := 0; i < len(nodeArr); i++ {
		NodeMap[nodeArr[i].Id] = &nodeArr[i]
	}
	return nodeArr
}

func StoreList() []StoreNode {
	return nodeArr
}

func StoreNodeGet(id int) StoreNode {
	return *NodeMap[id]
}

// 新增故事节点
func StoreNodeAdd(val string, input []int, output []int) int {
	var node StoreNode
	node.Id = 0
	node.Val = val
	node.Input = input
	node.Output = output
	for _, value := range nodeArr {
		if value.Id > node.Id {
			node.Id = value.Id
		}
	}
	node.Id += 1

	// 存储新节点
	NodeMap[node.Id] = &node
	nodeArr = append(nodeArr, node)
	storeSave()

	return node.Id
}

// 插入链接故事节点
func StoreNodeLink(val string, linkInput int, linkOutput int) {
	// 生成新节点
	storeId := StoreNodeAdd(val, []int{linkInput}, []int{linkOutput})
	// 断开旧链接
	NodeMap[linkInput].Output = append(NodeMap[linkInput].Output, storeId)
	NodeMap[linkOutput].Input = append(NodeMap[linkOutput].Input, storeId)
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
	storeSave()
}

func storeSave() {
	str, err := json.Marshal(nodeArr)
	if err != nil {
		fmt.Println("transfer err", err)
	}
	ioutil.WriteFile("file/example.json", str, 0644)
}
