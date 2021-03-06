package controller

// 建立用户表
type AuthTokenDetail struct {
	Token      string
	OwnRunning []AuthRunningNode // 当前所属团
}

type AuthRunningNode struct {
	RoomId int
	Role   string // kp, pc, 观察者
	// 操控角色
}

type AuthTokenTable struct {
	AuthTokenList []AuthTokenDetail
	AuthTokenMap  map[string]*AuthTokenDetail
}

var AuthTokenList []AuthTokenDetail
var AuthTokenMap map[string]*AuthTokenDetail

func AuthInit() {
	AuthTokenMap = make(map[string]*AuthTokenDetail)
	AuthTokenList = make([]AuthTokenDetail, 1)
	AuthTokenList[0] = AuthTokenDetail{
		Token: "test",
		OwnRunning: []AuthRunningNode{
			{
				RoomId: 1,
				Role:   "kp",
			},
			{
				RoomId: 2,
				Role:   "kp",
			},
		},
	}
	updateAuthTokenMap()
}

// token可用检测
func AuthCheck(token string, role string, RoomId int) bool {
	data, ok := AuthTokenMap[token]
	if !ok {
		return false
	} else {
		runningNode := getAuthRunning(data.OwnRunning, RoomId)
		if runningNode != nil {
			return runningNode.Role == role
		} else {
			return false
		}
	}
}

// 获取跑团故事详细
func getAuthRunning(arr []AuthRunningNode, roomId int) *AuthRunningNode {
	for i := 0; i < len(arr); i++ {
		if arr[i].RoomId == roomId {
			return &arr[i]
		}
	}
	return nil
}

// 回复登录状态
func AuthStatus(token string) *AuthTokenDetail {
	// 角色
	// 当前团
	// ip
	return AuthTokenMap[token]
}

func AuthKpNumGet(token string) int {
	auth := AuthStatus(token)
	num := 0
	for i := 0; i < len(auth.OwnRunning); i++ {
		if auth.OwnRunning[i].Role == "kp" {
			num += 1
		}
	}
	return num
}

func updateAuthTokenMap() {
	for i := 0; i < len(AuthTokenList); i++ {
		AuthTokenMap[AuthTokenList[i].Token] = &AuthTokenList[i]
	}
}
