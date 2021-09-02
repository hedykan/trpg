package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	con "github.com/trpg/controller"

	route "github.com/trpg/route"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func main() {
	port := ":12345"
	mainInit()
	fmt.Printf("server port%s\n", port)
	openHtml()
	http.ListenAndServe(port, nil)
	return
}

func mainInit() {
	con.StoryInit()
	con.RunInit()
	route.RouteInit()
	route.LogInit()
}

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func openHtml() {
	path, _ := os.Getwd()
	htmlPath := path + "\\dist\\index.html"
	_, err := os.Stat(htmlPath)
	if err == nil {
		open(htmlPath)
	} else {
		fmt.Printf("%s is not exist", htmlPath)
	}
}

func open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(`cmd`, `/c`, run, uri)
	return cmd.Start()
}
