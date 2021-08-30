package main

import (
	"net/http"

	con "github.com/trpg/controller"

	route "github.com/trpg/route"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func main() {
	port := ":12345"
	mainInit()
	http.ListenAndServe(port, nil)
	return
}

func mainInit() {
	con.StoryInit()
	con.RunInit()
	route.RouteInit()
}
