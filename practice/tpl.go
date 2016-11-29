package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", readDir) //设置访问路由

	err := http.ListenAndServe(":8080", nil) //设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func readDir(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/bootstrap/css/bootstrap.min.css" {
		http.ServeFile(w, r, "./bootstrap/css/bootstrap.min.css")
		return
	}

	if r.URL.Path == "/bootstrap/js/bootstrap.min.js" {
		http.ServeFile(w, r, "./bootstrap/js/bootstrap.min.js")
		return
	}

	//	fmt.Fprint(w, "Hello golang http!")

	//	files, _ := ioutil.ReadDir("./")

	//	for _, f := range files {
	//		fmt.Fprint(w, f.IsDir())
	//		fmt.Fprint(w, f.Mode())
	//		fmt.Fprint(w, f.ModTime())
	//		fmt.Fprint(w, f.Name())
	//		fmt.Fprint(w, f.Size())

	//		fmt.Fprint(w, "\n")
	//	}
	t, err := template.ParseFiles("tpl.html")
	t.Execute(w, "")
	if err != nil {
		log.Fatal(err)
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
