package model

import "os"

func dirCreate(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		os.MkdirAll(dir, 0755)
	}
}
