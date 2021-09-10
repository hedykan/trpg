package controller

import (
	"encoding/json"
	"io/ioutil"
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
	f, err := ioutil.ReadFile("file/attr_example.json")
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

func attrArrSave(attrArr []AttrNode) {
	str, err := json.Marshal(attrArr)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("file/attr_example.json", str, 0644)
}

func updateAttrNodeMap() {
	for i := 0; i < len(AttrNodeArr); i++ {
		AttrNodeMap[AttrNodeArr[i].Id] = &AttrNodeArr[i]
	}
}
