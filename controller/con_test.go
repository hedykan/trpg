package controller

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	AuthInit()
	num := AuthKpNumGet("test")
	fmt.Println(num)
}
