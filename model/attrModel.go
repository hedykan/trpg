package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type AttrNode struct {
	Id  int
	Val string
	Num int
}

func AttrLoad(addr string) []AttrNode {
	var table []AttrNode
	f, err := ioutil.ReadFile(addr)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &table)
	if err != nil {
		panic(err)
	}
	return table
}

func AttrSave(attrNodeList []AttrNode, addr string) {
	strArr := strings.Split(addr, "/")
	dir := strings.Join(strArr[:len(strArr)-1], "/")
	dirCreate(dir)
	str, err := json.Marshal(attrNodeList)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(addr, str, 0644)
}
