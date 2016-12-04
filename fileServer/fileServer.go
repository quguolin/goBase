package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {

	files, _ := ioutil.ReadDir("./")

	file_infs := make(map[int]map[string]string)

	for index, f := range files {
		file_inf := make(map[string]string)

		file_inf["is_dir"] = strconv.FormatBool(f.IsDir())
		file_inf["name"] = f.Name()
		fmt.Println(f.Size())
		file_inf["size"] = strconv.FormatInt(f.Size(), 10)

		t := f.ModTime()
		file_inf["mod_time"] = t.Format("2006-01-02 15:04:05")

		file_inf["mode"] = strconv.FormatInt(f.Size(), 10)

		file_infs[index] = file_inf
	}

	for _, value := range file_infs {
		fmt.Println(value)
	}

	//	fmt.Println(test)
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
