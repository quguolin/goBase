package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	files, _ := ioutil.ReadDir("./")

	for _, f := range files {
		file_info := make(map[string]string)
		//		file_info["is_dir"] = string(f.IsDir())
		//		file_info["mod_tile"] = string(f.ModTime())

		file_info["name"] = string(f.Name())
		file_info["size"] = string(f.Size())

		fmt.Println(file_info)
	}

	//	http.HandleFunc("/", readDir) //设置访问路由

	//	err := http.ListenAndServe(":8080", nil) //设置监听端口
	//	if err != nil {
	//		log.Fatal("ListenAndServe:", err)
	//	}
}

type file struct {
	Name  string
	IsDir bool
	Size  uint32
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

	//	file_info = make(map[string]string)
	//	files_info = make(map[int]map[string]string)

	//	files, _ := ioutil.ReadDir("./")
	//	for _, file := range files {

	//	}

	//	t, err := template.ParseFiles("fileServer.html")
	//	t.Execute(w, dir)
	//	if err != nil {
	//		log.Fatal(err)
	//	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
