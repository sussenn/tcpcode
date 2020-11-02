package main

import (
	"fmt"
	"net/http"
)

//request 请求url/参数/头/行/体 的获取
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "请求url的获取: ", r.URL.Path)
	//如:/findById?id=xxx 获取到"id=xxx"
	fmt.Fprintln(w, "请求的参数的获取: ", r.URL.RawQuery)
	//请求头所有信息的获取
	fmt.Fprintln(w, "请求头的获取: ", r.Header)
	//返回结果是切片
	fmt.Fprintln(w, "请求头指定参数信息的获取: ", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头指定参数属性的获取: ", r.Header.Get("Accept-Encoding"))
	//------------------------------------------------------
	//请求体的获取
	//先获取请求体内容长度
	//len := r.ContentLength
	////用于接收请求体内容的切片
	//body := make([]byte, len)
	//r.Body.Read(body)
	//fmt.Fprintln(w, "请求体内容: ", string(body))
	//---------------------------------------------------
	//form表单数据的获取	[[注意:请求体的数据在上面r.Body.Read(body)读取后,将无法再次被获得]]
	//必须先调用此函数,解析表单
	r.ParseForm()
	//fmt.Fprintln(w, "POST表单内容: ", r.Form.Get("username"))
	//fmt.Fprintln(w, "POST表单内容: ", r.PostForm.Get("username"))
	fmt.Fprintln(w, "POST表单内容: ", r.PostFormValue("username"))

}

func main() {
	http.HandleFunc("/test", myHandler)
	http.ListenAndServe(":8080", nil)
}
