package route

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	con "github.com/trpg/controller"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Local(), r.Method, r.URL)
		log.Println(r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
func test2Middle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func methodMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// 解决跨域问题
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		switch r.Method {
		// 复杂POST处理
		case "OPTIONS":
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := getToken(r)
		roomId := getRoomId(r)
		check := con.AuthCheck(token, "kp", roomId)
		if check {
			next.ServeHTTP(w, r)
		} else {
			errResInput(w, r, 201)
			return
		}
	})
}

// 新增kp, pc, 观察者身份检测
func middleware(next func(http.ResponseWriter, *http.Request)) http.Handler {
	return logMiddleware(methodMiddleware(http.HandlerFunc(next)))
}

func checkMiddleware(next func(http.ResponseWriter, *http.Request)) http.Handler {
	return logMiddleware(methodMiddleware(authMiddleware(http.HandlerFunc(next))))
}

func LogInit() {
	file := "./logs/http.log"
	_, err := os.Stat("./logs")
	if err != nil {
		os.MkdirAll("./logs", 0755)
	}
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

func getRoomId(r *http.Request) int {
	switch r.Method {
	case "GET":
		query := get(r)
		roomId, _ := strconv.Atoi(query["roomId"])
		return roomId
	case "POST":
		var query struct{ RoomId int }
		postJson(r, &query)
		roomId := query.RoomId
		return roomId
	default:
		return -1
	}
}
