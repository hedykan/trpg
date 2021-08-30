package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	con "github.com/trpg/controller"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	con.StoryCreate()
	con.RunStatusCreate()
	res := res(nil)
	enc.Encode(res)
}

func storyList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	res := res(con.StoryList())
	enc.Encode(res)
}

func storyGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	query := get(r)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		fmt.Println("transfe err", err)
	}
	res := res(con.StoryNodeGet(id))
	enc.Encode(res)
}

func runStatusList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	res := res(con.RunStatusList())
	enc.Encode(res)
}

func runStep(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	query := get(r)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		fmt.Println("transfe err", err)
	}
	con.RunStep(id)
	res := res(nil)
	enc.Encode(res)
}
