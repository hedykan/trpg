package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type RunVoteStatus struct {
	SelecterId int
	Num        int
}

type RunVote struct {
	NodeId         int // 当前节点i
	VoteStatusList []RunVoteStatus
	TokenList      []string
}

type RunStatus struct {
	NowStoryNode    int
	RecordStoryNode []int
	RecordVote      []RunVote
}

func StatusLoad(addr string) RunStatus {
	var res RunStatus
	f, err := ioutil.ReadFile(addr)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &res)
	if err != nil {
		panic(err)
	}
	return res
}

func StatusSave(status RunStatus, addr string) {
	strArr := strings.Split(addr, "/")
	dir := strings.Join(strArr[:len(strArr)-1], "/")
	dirCreate(dir)
	str, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(addr, str, 0644)
}
