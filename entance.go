package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	//http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/", myWeb)

	//http.Handle("/js/", http.FileServer(http.Dir("./js")))

	//http.Handle("/js/",http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))

	////指定相对路径./static 为文件服务路径
	//staticHandle := http.FileServer(http.Dir("./static"))
	////将/js/路径下的请求匹配到 ./static/js/下
	//http.Handle("/js/", staticHandle)

	fmt.Println("服务器即将开启，访问地址 http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("服务器开启错误")
	}
}

func myWeb(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "这是一个开始")
	for k, v := range r.PostForm {
		fmt.Println("key : ", k, "value : ", v)
	}
	for k, v := range r.URL.Query() {
		fmt.Println("key : ", k, "value : ", v)
	}

	//t := template.New("index")
	//t.Parse("<div id='templateTextDiv'>Hi,{{.name}},{{.someStr}}</div>")
	t, _ := template.ParseFiles("./static/index.html")
	data := map[string]string{
		"name": "aaa",
		"say":  "ღ( ´･ᴗ･` )比心",
	}
	query := r.URL.Query()
	get := query.Get("name")
	if get != "" {
		data["name"] = get
	}
	t.Execute(w, data)

}
