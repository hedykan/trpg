package route

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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

func middleware(next func(http.ResponseWriter, *http.Request)) http.Handler {
	return logMiddleware(http.HandlerFunc(next))
}

func LogInit() {
	file := "./logs/http.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
