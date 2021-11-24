package controller

// 角色
type RoleNode struct {
	Id          int
	Name        string
	AttrList    []AttrNode
	RoleAttrMap map[int]*AttrNode
}

// 角色拥有状态节点
type RoleStatusNode struct {
}

type RoleTable struct {
	RoleList []RoleNode
	RoleMap  map[int]*RoleNode
}

type RoleListType []RoleNode

func RoleCreate(name string, list []AttrNode) RoleNode {
	var node RoleNode
	node.AttrList = list
	node.updateMap()
	return node
}

func RoleNodeGet(table RoleTable, RoleId int) RoleNode {
	return *table.RoleMap[RoleId]
}

func RoleNodeAdd(table RoleTable, name string, list []AttrNode) {
	node := RoleCreate(name, list)
	node.Id = RoleListType(table.RoleList).getMaxId() + 1
	table.RoleList = append(table.RoleList, node)
	table.updateMap()
}

func RoleAttrOperate(role RoleNode, attrId int, operate string, num int) {
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
