package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request)  {
	// 定义函数
	kua := func(arg string) (string, error) {
		return arg + "真帅", nil
	}

	// 定义模版
	// 解析模版
	t := template.New("f.tmpl")   // 名字一定要和模版的名字都能对上
	// 告诉模版引擎，我多了一个自定义函数kua
	t.Funcs(template.FuncMap{
		"kua": kua,
	})
	_,err := t.ParseFiles("./f.tmpl")

	if err != nil{
		fmt.Println("错误")
		return
	}

	// 渲染模版
	name := "小王子"
	t.Execute(w, name)

}

func main()  {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err %v", err)
		return
	}
}
