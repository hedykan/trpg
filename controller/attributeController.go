package controller

import (
	"encoding/json"
	"io/ioutil"
)

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

func updateAttrNodeMap() {
	for i := 0; i < len(AttrNodeArr); i++ {
		AttrNodeMap[AttrNodeArr[i].Id] = &AttrNodeArr[i]
	}
}

func AttrInit() {
	AttrNodeArr = append(AttrNodeArr, AttrNodeCreate("test", 0))
	attrArrSave(AttrNodeArr)
}

func attrArrSave(attrArr []AttrNode) {
	str, err := json.Marshal(attrArr)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("file/attr_example.json", str, 0644)
}
