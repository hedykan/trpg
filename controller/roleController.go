package controller

import (
	"fmt"

	"github.com/trpg/model"
)

// 角色
type RoleNode struct {
	Id          int
	Name        string
	AttrList    []AttrNode
	RoleAttrMap map[int]*AttrNode
}

type RoleTable struct {
	RoleList []RoleNode
	RoleMap  map[int]*RoleNode
}

type RoleListType []RoleNode

var RoleTestTable RoleTable

func RoleTest() {
	RoleTestTable.RoleList = RoleArrTransfer(model.RoleLoad("./file/room/1/" + "role.json"))
	RoleTestTable.RoleMap = make(map[int]*RoleNode)
	RoleTestTable.updateMap()
	// RoleNodeAdd(&RoleTestTable, "test1", RoomMap[1].Attribute.AttrList)
	fmt.Println(RoleTestTable)
	model.RoleSave(RoleArrTransferModel(RoleTestTable.RoleList), "./file/room/1/"+"role.json")
}

func RoleCreate(name string, list []AttrNode) *RoleNode {
	var node RoleNode
	node.AttrList = list
	node.RoleAttrMap = make(map[int]*AttrNode)
	node.updateMap()
	return &node
}

func RoleNodeGet(table *RoleTable, RoleId int) RoleNode {
	return *table.RoleMap[RoleId]
}

func RoleNodeAdd(table *RoleTable, name string, list []AttrNode) {
	node := RoleCreate(name, list)
	node.Id = RoleListType(table.RoleList).getMaxId() + 1
	fmt.Println(node.Id)
	table.RoleList = append(table.RoleList, *node)
	table.updateMap()
}

// 角色属性操作
func RoleAttrOperate(role *RoleNode, attrId int, operate string, num int) {
	switch operate {
	case "add":
		role.RoleAttrMap[attrId].Num += num
		break
	case "sub":
		role.RoleAttrMap[attrId].Num -= num
		break
	}
}

func (roleNode *RoleNode) updateMap() {
	for i := 0; i < len(roleNode.AttrList); i++ {
		roleNode.RoleAttrMap[roleNode.AttrList[i].Id] = &roleNode.AttrList[i]
	}
}

func (table *RoleTable) updateMap() {
	for i := 0; i < len(table.RoleList); i++ {
		table.RoleMap[table.RoleList[i].Id] = &table.RoleList[i]
	}
}

func (roleList RoleListType) getMaxId() int {
	max := 0
	for i := 0; i < len(roleList); i++ {
		if max < roleList[i].Id {
			max = i
		}
	}
	return max
}
