package main



import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Gender string
	Age int
}

func sayhello(w http.ResponseWriter, r *http.Request)  {
	// 解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("解析模版 failed, err %v", err)
		return
	}
	// 渲染模版
	// 1. 传递结构体
	u1 := User{
		"小王子",
		"男",
		18,
	}
	//t.Execute(w, u1)

	// 2. 传递map
	m1 := map[string]interface{}{
		"name": "小王子",
		"gender": "男",
		"age": 18,
	}
	//t.Execute(w, m1)

	// 3. 传递结构体和map
	hobbyList := []string{"篮球", "足球", "双色球"}
	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
		"hobby": hobbyList,
	})

}

func main()  {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err %v", err)
		return
	}
}
