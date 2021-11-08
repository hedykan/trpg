package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

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

func StoryLoad(addr string) []StoryNode {
	var res []StoryNode
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

// 保存故事
func StorySave(arr []StoryNode, addr string) {
	strArr := strings.Split(addr, "/")
	dir := strings.Join(strArr[:len(strArr)-1], "/")
	dirCreate(dir)
	str, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(addr, str, 0644)
}

func StoryBackgroundLoad(addr string) StoryBackground {
	var res StoryBackground
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

func StoryBackgroundSave(storyBackground StoryBackground, addr string) {
	strArr := strings.Split(addr, "/")
	dir := strings.Join(strArr[:len(strArr)-1], "/")
	dirCreate(dir)
	str, err := json.Marshal(storyBackground)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(addr, str, 0644)
}
