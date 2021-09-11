package controller

import (
	"os"
	"strings"
)

func FileCheck(path string, function func()) {
	_, err := os.Stat(path)
	if err != nil {
		strArr := strings.Split(path, "/")
		dir := strings.Join(strArr[:len(strArr)-1], "/")
		DirCheck(dir)
		function()
	}
}

func DirCheck(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		os.MkdirAll(dir, 0755)
	}
}
