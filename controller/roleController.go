package controller

import (
	"strconv"

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

func RoleTableCreate() *RoleTable {
	var table RoleTable
	table.RoleList = make([]RoleNode, 0)
	table.updateMap()
	return &table
}

// 获取角色节点
func RoleNodeGet(table *RoleTable, RoleId int) RoleNode {
	return *table.RoleMap[RoleId]
}

// 获取角色列表
func RoleNodeList(table *RoleTable) []RoleNode {
	return table.RoleList
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
	table.updateMap()
}

// 删除角色节点
func RoleNodeDelete(table *RoleTable, roleId int) {
	table.deleteNode(roleId)
}

// 角色属性操作
func RoleAttrOperate(table *RoleTable, roleId int, attrId int, operate string, num int) {
	switch operate {
	case "add":
		table.RoleMap[roleId].RoleAttrMap[attrId].Num += num
		break
	case "sub":
		table.RoleMap[roleId].RoleAttrMap[attrId].Num -= num
		break
	}
}

func (table *RoleTable) save(roomId int) {
	addr := "./file/room/" + strconv.Itoa(roomId) + "/role.json"
	model.RoleSave(RoleArrTransferModel(table.RoleList), addr)
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

func (table *RoleTable) deleteNode(nodeId int) {
	for i := 0; i < len(table.RoleList); i++ {
		if nodeId == table.RoleList[i].Id {
			table.RoleList = append(table.RoleList[:i], table.RoleList[i+1:]...)
			break
		}
	}
	delete(table.RoleMap, nodeId)
}
