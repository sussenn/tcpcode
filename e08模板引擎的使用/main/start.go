package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/testTemplate", myTemplateHandler)
	http.HandleFunc("/testTemplate2", myTemplateHandler2)
	http.ListenAndServe(":2101", nil)
}

//模板引擎的设置
func myTemplateHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板文件	P1:文件所在路径2种写法都可以
	//t, err := template.ParseFiles("e08模板引擎的使用\\main\\index.html")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//Must()函数帮忙处理了err
	t := template.Must(template.ParseFiles("e08模板引擎的使用/main/index.html"))
	//执行
	//p2: 传入数据,页面使用{{.}}进行获取
	t.Execute(w, "hello template")
}

func myTemplateHandler2(w http.ResponseWriter, r *http.Request) {
	//可传入多个文件,但默认输出到第一个	p: 文件相对路径
	t, err := template.ParseFiles("e08模板引擎的使用/main/index.html", "e08模板引擎的使用/main/hello.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	//可以指定传输文件的路径	p2: 文件名
	s := t.ExecuteTemplate(w, "hello.html", "指定传输到hello.html文件")
	fmt.Println(s)
}
