package main

import (
	"fmt"
	"net/http"

	con "github.com/trpg/controller"

	route "github.com/trpg/route"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func main() {
	port := ":12345"
	mainInit()
	fmt.Printf("server port%s\n", port)
	http.ListenAndServe(port, nil)
	return
}

func mainInit() {
	con.StoryInit()
	con.RunInit()
	route.RouteInit()
}
