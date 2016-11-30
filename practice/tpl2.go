package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", temp) //设置访问路由

	err := http.ListenAndServe(":8080", nil) //设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func temp(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./tpl2.html")
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Title string
	}{
		Title: "golang html template",
	}

	//	err = t.Execute(os.Stdout, data)
	err = t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
