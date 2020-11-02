package main

import (
	"fmt"
	"net/http"
)

//自定义处理器函数(入参固定写法)
func myHandler(w http.ResponseWriter, r *http.Request) {
	//p2:页面输出内容 p3:请求url
	fmt.Fprintln(w, "hello world", r.URL.Path)
}
func main() {
	//配置适配器 	p1:请求路径 p2:自定义的处理器
	//http.HandleFunc()方法会自动处理入参自定义的处理器
	http.HandleFunc("/", myHandler)
	//创建路由 	p1:监听的端口号 p2:多路复用器(相当于Java的前端控制器),nil使用默认的
	http.ListenAndServe(":8080", nil)
}
