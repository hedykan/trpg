package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func dirCreate(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		os.MkdirAll(dir, 0755)
	}
}

func fileSave(data interface{}, addr string) {
	strArr := strings.Split(addr, "/")
	dir := strings.Join(strArr[:len(strArr)-1], "/")
	dirCreate(dir)
	str, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(addr, str, 0644)
}
