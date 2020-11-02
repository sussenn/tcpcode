package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "自行配置server进行监听")
}
func main() {
	myHandler := MyHandler{}
	//配置server		访问地址默认"\"
	server := http.Server{
		Addr:        ":3111",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second, //2s超时
	}
	//使用已配置的server进行监听
	server.ListenAndServe()
}
