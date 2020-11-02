package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "使用自定义多路复用器发送请求")
}
func main() {
	//创建多路复用器
	mux := http.NewServeMux()
	mux.HandleFunc("/mux", myHandler)
	http.ListenAndServe(":8080", mux)
}
