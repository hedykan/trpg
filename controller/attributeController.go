package controller

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

/*
	建立一个属性表
*/

type AttrNode struct {
	Id  int
	Val string
	Num int
}

var AttrNodeArr []AttrNode
var AttrNodeMap map[int]*AttrNode

func AttrNodeCreate(val string, num int) AttrNode {
	node := AttrNode{Id: 0, Val: val, Num: num}
	for _, value := range AttrNodeArr {
		if value.Id > node.Id {
			node.Id = value.Id
		}
	}
	node.Id += 1
	return node
}

func AttrCreate() {
	AttrNodeArr = append(AttrNodeArr, AttrNodeCreate("test", 0))
	attrArrSave(AttrNodeArr)
}

func AttrLoad() []AttrNode {
	AttrNodeMap = make(map[int]*AttrNode)
	f, err := ioutil.ReadFile("./file/attr_example.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &AttrNodeArr)
	if err != nil {
		panic(err)
	}
	updateAttrNodeMap()

	return AttrNodeArr
}

// 属性列表初始化
func AttrInit() {
	path, _ := os.Getwd()
	path = path + "/file/attr_example.json"
	FileCheck(path, AttrCreate)
	AttrLoad()
	updateAttrNodeMap()
}

// 获取属性列表
func AttrList() []AttrNode {
	return AttrNodeArr
}

// 获取单一属性
func AttrNodeGet(id int) AttrNode {
	return *AttrNodeMap[id]
}

// 新增属性节点
func AttrNodeAdd(val string, num int) int {
	node := AttrNodeCreate(val, num)
	AttrNodeArr = append(AttrNodeArr, node)
	updateAttrNodeMap()
	return node.Id
}

// 编辑属性节点
func AttrNodeEdit(id int, val string, num int) bool {
	if _, ok := AttrNodeMap[id]; ok != true {
		return false
	}
	AttrNodeMap[id].Val = val
	AttrNodeMap[id].Num = num
	return true
}

func AttrNodeNumAdd(id int, num int) bool {
	if data, ok := AttrNodeMap[id]; ok != true {
		return false
	} else {
		data.Num += num
		return true
	}
}

// 删除属性节点
func AttrNodeDelete(id int) bool {
	if _, ok := AttrNodeMap[id]; ok != true {
		return false
	}
	// TODO 删除操作
	delete(AttrNodeMap, id)
	index := searchAttrId(AttrNodeArr, id)
	if index == -1 {
		return false
	}
	deleteAttrSlice(AttrNodeArr, index)

	return true
}

// 保存属性节点
func attrArrSave(attrArr []AttrNode) {
	str, err := json.Marshal(attrArr)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("file/attr_example.json", str, 0644)
}

func deleteAttrSlice(arr []AttrNode, index int) []AttrNode {
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func searchAttrId(arr []AttrNode, id int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i].Id == id {
			return i
		}
	}
	return -1
}

func updateAttrNodeMap() {
	for i := 0; i < len(AttrNodeArr); i++ {
		AttrNodeMap[AttrNodeArr[i].Id] = &AttrNodeArr[i]
	}
}
