package controller

import (
	"fmt"
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

func RoleTableCreate() *RoleTable {
	var table RoleTable
	table.RoleList = make([]RoleNode, 0)
	table.updateMap()
	return &table
}

// func RoleCreate(name string, list []AttrNode) *RoleNode {
// 	var node RoleNode
// 	node.AttrList = list
// 	node.RoleAttrMap = make(map[int]*AttrNode)
// 	node.updateMap()
// 	return &node
// }

// 获取角色节点
func RoleNodeGet(table *RoleTable, RoleId int) RoleNode {
	return *table.RoleMap[RoleId]
}

// 新增角色节点
func RoleNodeAdd(table *RoleTable, name string, list []AttrNode) {
	node := &RoleNode{
		Id:          RoleListType(table.RoleList).getMaxId() + 1,
		Name:        name,
		AttrList:    list,
		RoleAttrMap: map[int]*AttrNode{},
	}
	node.updateMap()
	table.RoleList = append(table.RoleList, *node)
	fmt.Println(table.RoleList)
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
			max = roleList[i].Id
		}
	}
	return max
}
