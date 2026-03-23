package main

import (
	"fmt"
	"net/http"
	"os" // 新增
)

// 见鬼了，这个文件是干嘛的？
func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	// 新增接口
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		app := os.Getenv("APP_NAME")
		if app == "" {
			app = "default-app"
		}
		fmt.Fprintf(w, "app_name=%s", app)
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("HTTP server listening on :8080...")
}
