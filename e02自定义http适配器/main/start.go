package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

//自定义适配器	实现ServeHTTP(w http.ResponseWriter, r *http.Request)方法
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自定义适配器处理请求")
}
func main() {
	myHandler := MyHandler{}
	//使用自定义适配器
	//http.Handle()需要入参实现ServeHTTP()方法
	http.Handle("/myHandler", &myHandler)
	http.ListenAndServe(":3110", nil)
}
