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
	// con.StoreNodeAdd("test", []int{1, 2}, []int{2})
	con.StoreNodeLink("link test", 1, 2)
	res := res(nil)
	enc.Encode(res)
}

func storeList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	res := res(con.StoreList())
	enc.Encode(res)
}

func storeGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	query := get(r)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		fmt.Println("transfe err", err)
	}
	res := res(con.StoreNodeGet(id))
	enc.Encode(res)
}
