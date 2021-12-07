package controller

/*
	底层思想是操作一个故事表，根据输入输出节点决定故事走向
*/
// TODO 整合故事与背景结构体
// 以storyNodeArr/storyNodeMap为操作实体
// 新的保存方式

type StoryTable struct {
	StoryList []StoryNode
	StoryMap  map[int]*StoryNode
}

type StoryListType []StoryNode

// 输入输出结构体
// 可以添加ext扩展，使run控制器根据条件显示故事选项
type StorySeleter struct {
	Id  int
	Val string
}

type StoryNode struct {
	Id     int
	Val    string
	Input  []StorySeleter // []int
	Output []StorySeleter // []int
}

type StoryBackground struct {
	Background string
}

func StoryBackgroundEdit(backgroundTable *StoryBackground, background string) {
	backgroundTable.Background = background
}

// 改为回storytable
func StoryCreate() []StoryNode {
	storyArr := make([]StoryNode, 2)
	storyArr[0] = StoryNode{
		Id:     0,
		Input:  []StorySeleter{},
		Output: []StorySeleter{{Id: 1, Val: "end"}},
		Val:    "start",
	}
	storyArr[1] = StoryNode{
		Id:     1,
		Input:  []StorySeleter{{Id: 0, Val: "start"}},
		Output: []StorySeleter{},
		Val:    "end",
	}
	return storyArr
}

// 故事整体展示
func StoryList(storyNodeArr []StoryNode) []StoryNode {
	return storyNodeArr
}

// 故事节点获取
func StoryNodeGet(table *StoryTable, id int) StoryNode {
	if data, ok := table.StoryMap[id]; ok {
		return *data
	} else {
		return StoryNode{}
	}
}

func StoryNodeCreate(storyNodeArr []StoryNode, val string) StoryNode {
	var node = StoryNode{
		Id:     0,
		Input:  nil,
		Output: nil,
		Val:    val,
	}
	for _, value := range storyNodeArr {
		if value.Id > node.Id {
			node.Id = value.Id
		}
	}
	node.Id += 1

	return node
}

// 新增故事节点 TODO 输入输出组节点非必须
func StoryNodeAdd(table *StoryTable, val string, input []StorySeleter, output []StorySeleter) bool {
	node := StoryNodeCreate(table.StoryList, val)
	// append后StoryNodeMap的地址和append的地址不同, 需要更新StoryNodeMap
	table.StoryList = append(table.StoryList, node)
	// updateNodeMap(table)
	table.updateMap()
	// 自己的输出组
	for i := 0; i < len(output); i++ {
		StorySelecterAdd(table, node.Id, output[i].Id, output[i].Val)
	}
	// 对方的输出组
	for i := 0; i < len(input); i++ {
		StorySelecterAdd(table, input[i].Id, node.Id, input[i].Val)
	}

	return true
}

// 故事节点修改
// TODO 选择组非必须
// 更新投票节点
// 如果节点已跑，输出节点不可改
func StoryNodeEdit(table *StoryTable, nodeId int, val string, input []StorySeleter, output []StorySeleter) bool {
	if data, ok := table.StoryMap[nodeId]; !ok {
		return false
	} else {
		// 先清掉链接
		// 清除自己的输出组
		for i := 0; i < len(data.Output); i++ {
			StorySelecterDelete(table, nodeId, data.Output[i].Id)
		}
		// 清除别人的输出组
		for i := 0; i < len(data.Input); i++ {
			StorySelecterDelete(table, data.Input[i].Id, nodeId)
		}
	}
	// 设置自己的输出组
	for i := 0; i < len(output); i++ {
		StorySelecterAdd(table, nodeId, output[i].Id, output[i].Val)
	}
	// 对方的输出组
	for i := 0; i < len(input); i++ {
		StorySelecterAdd(table, input[i].Id, nodeId, input[i].Val)
	}
	table.StoryMap[nodeId].Val = val
	return true
}

// 故事节点删除
func StoryNodeDelete(table *StoryTable, nodeId int) {
	if nodeId == 1 || nodeId == 0 {
		return
	}
	data, ok := table.StoryMap[nodeId]
	if ok {
		var index int
		// 清除自己的输出组
		for i := 0; i < len(data.Output); i++ {
			StorySelecterDelete(table, nodeId, data.Output[i].Id)
		}
		// 清除别人的输出组
		for i := 0; i < len(data.Input); i++ {
			StorySelecterDelete(table, data.Input[i].Id, nodeId)
		}
		delete(table.StoryMap, nodeId)
		index = searchStoryId(table.StoryList, nodeId)
		table.StoryList = deleteNodeSlice(table.StoryList, index)
	}
}

// 故事节点链接添加
func StorySelecterAdd(table *StoryTable, nodeId int, linkId int, val string) bool {
	// 新增输出节点
	if node, ok := table.StoryMap[nodeId]; ok {
		index := searchSelecterId(node.Output, linkId)
		if index == -1 {
			node.Output = append(node.Output, StorySeleter{Id: linkId, Val: val})
		} else {
			node.Output[index].Val = val
		}
	} else {
		return false
	}
	// 被输入的节点新增输入节点
	if node, ok := table.StoryMap[linkId]; ok {
		index := searchSelecterId(node.Input, nodeId)
		if index == -1 {
			node.Input = append(node.Input, StorySeleter{Id: nodeId, Val: val})
		} else {
			node.Input[index].Val = val
		}
	} else {
		return false
	}

	return true
}

// 故事节点链接删除
// 查找自己的输出组删除对方输入组
// 查找自己的输入组让对方删除自己
func StorySelecterDelete(table *StoryTable, nodeId int, linkId int) bool {
	// 删除输出组链接
	if node, ok := table.StoryMap[nodeId]; ok {
		index := searchSelecterId(node.Output, linkId)
		if index != -1 {
			node.Output = deleteSelecter(node.Output, index)
		}
	} else {
		return false
	}
	// 删除输入组链接
	if node, ok := table.StoryMap[linkId]; ok {
		index := searchSelecterId(node.Input, nodeId)
		if index != -1 {
			node.Input = deleteSelecter(node.Input, index)
		}
	} else {
		return false
	}

	return true
}

// 检查输入输出节点是否存在
func storyCheckSelecter(arr []StorySeleter, arrMap map[int]*StoryNode) bool {
	for _, v := range arr {
		_, ok := arrMap[v.Id]
		if !ok {
			return false
		}
	}
	return true
}

func searchSelecterId(idArr []StorySeleter, id int) int {
	for i := 0; i < len(idArr); i++ {
		if id == idArr[i].Id {
			return i
		}
	}
	return -1
}

func searchId(idArr []int, id int) int {
	for i := 0; i < len(idArr); i++ {
		if id == idArr[i] {
			return i
		}
	}
	return -1
}

func searchIdReverse(idArr []int, id int) int {
	for i := len(idArr) - 1; i >= 0; i-- {
		if id == idArr[i] {
			return i
		}
	}
	return -1
}

func searchStoryId(arr []StoryNode, id int) int {
	for i := 0; i < len(arr); i++ {
		if id == arr[i].Id {
			return i
		}
	}
	return -1
}

func deleteSelecter(arr []StorySeleter, index int) []StorySeleter {
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func deleteNodeSlice(arr []StoryNode, index int) []StoryNode {
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

// func updateNodeMap(table *StoryTable) {
// 	for i := 0; i < len(table.StoryList); i++ {
// 		table.StoryMap[table.StoryList[i].Id] = &table.StoryList[i]
// 	}
// }

func (storyTable *StoryTable) updateMap() {
	for i := 0; i < len(storyTable.StoryList); i++ {
		storyTable.StoryMap[storyTable.StoryList[i].Id] = &storyTable.StoryList[i]
	}
}
