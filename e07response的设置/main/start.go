package main

import (
	"encoding/json"
	"net/http"
	"tcpcode/e07response的设置/model"
)

func main() {
	//测试json
	http.HandleFunc("/testJson", myResJsonHandler)
	//测试重定向
	http.HandleFunc("/testRedirect", myResRedirectHandler)
	http.ListenAndServe(":2101", nil)
}

//响应json
func myResJsonHandler(w http.ResponseWriter, r *http.Request) {
	//设置响应json格式
	w.Header().Set("Content-Type", "application/json")
	u := model.User{
		ID:       1,
		Username: "zsc",
		Sex:      "1",
		Password: "123",
		Email:    "zsc@qq.com",
	}
	//将对象转json
	uJson, _ := json.Marshal(u)
	//响应
	w.Write(uJson)
}

//重定向设置
func myResRedirectHandler(w http.ResponseWriter, r *http.Request) {
	//必须先设置 Location
	w.Header().Set("Location", "https://www.baidu.com")
	//再设置响应状态码
	w.WriteHeader(302)
}
