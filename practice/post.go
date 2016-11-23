package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	data := make(url.Values)

	data["key"] = []string{"this is key"}
	data["value"] = []string{"this is value"}

	res, err := http.PostForm("http://127.0.0.1:8080", data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer res.Body.Close()

	fmt.Println("post success")
}
