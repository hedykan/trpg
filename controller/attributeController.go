package controller

/*
	建立一个属性表
*/

type AttrNode struct {
	Id  int
	Val string
	Num int
}

type AttrTable struct {
	AttrList []AttrNode
	AttrMap  map[int]*AttrNode
}

type AttrListType []AttrNode

var AttrNodeArr []AttrNode
var AttrNodeMap map[int]*AttrNode

// 新建属性
func AttrNodeCreate(val string, num int) AttrNode {
	node := AttrNode{Id: 0, Val: val, Num: num}
	return node
}

func AttrCreate() *AttrTable {
	var attrTable AttrTable
	attrTable.AttrMap = make(map[int]*AttrNode)
	return &attrTable
}

// 获取属性列表
func AttrList(table *AttrTable) []AttrNode {
	return table.AttrList
}

// 获取单一属性
func AttrNodeGet(table *AttrTable, id int) AttrNode {
	return *table.AttrMap[id]
}

// 属性节点新增
func AttrNodeAdd(table *AttrTable, val string, num int) {
	node := AttrNodeCreate(val, num)
	for _, v := range table.AttrList {
		if v.Id > node.Id {
			node.Id = v.Id
		}
	}
	node.Id += 1
	// updateAttrNodeMap(table)
	table.updateMap()
	table.AttrList = append(table.AttrList, node)
}

// 编辑属性节点
func AttrNodeEdit(table *AttrTable, id int, val string, num int) bool {
	if data, ok := table.AttrMap[id]; ok {
		data.Val = val
		data.Num = num
		return true
	} else {
		return false
	}
}

func AttrNodeNumAdd(table *AttrTable, id int, num int) bool {
	if data, ok := table.AttrMap[id]; ok {
		data.Num += num
		return true
	} else {
		return false
	}
}

// 删除属性节点
func AttrNodeDelete(table *AttrTable, id int) bool {
	if _, ok := table.AttrMap[id]; !ok {
		return false
	}
	// TODO 删除操作
	delete(table.AttrMap, id)
	index := searchAttrId(table.AttrList, id)
	if index == -1 {
		return false
	}
	deleteAttrSlice(table.AttrList, index)

	return true
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

// func updateAttrNodeMap(attrTable *AttrTable) {
// 	for i := 0; i < len(attrTable.AttrList); i++ {
// 		attrTable.AttrMap[attrTable.AttrList[i].Id] = &attrTable.AttrList[i]
// 	}
// }

func (attrTable *AttrTable) updateMap() {
	for i := 0; i < len(attrTable.AttrList); i++ {
		attrTable.AttrMap[attrTable.AttrList[i].Id] = &attrTable.AttrList[i]
	}
}
