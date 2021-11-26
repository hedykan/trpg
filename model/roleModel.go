package model

import (
	"encoding/json"
	"io/ioutil"
)

type RoleNode struct {
	Id       int
	Name     string
	AttrList []AttrNode
}

func RoleLoad(addr string) []RoleNode {
	var res []RoleNode
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

func RoleSave(arr []RoleNode, addr string) {
	fileSave(arr, addr)
}
