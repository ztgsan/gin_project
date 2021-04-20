package main

import (
	"fmt"
	"html/template"
	"net/http"
)


func sayhello(w http.ResponseWriter, r *http.Request) {
	// 解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil{
		fmt.Printf("parse faild %v", err)
		return
	}
	// 渲染模版
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("err &v", err)
	}

}

func main()  {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err %v", err)
		return
	}
}


