package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//	"strconv"
)

func fGet(url string) {
	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	re, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(re))
	ch <- struct{}{}
}

var ch = make(chan struct{})

func main() {
	url := "mysqlCurrent.php"

	for index := 1; index <= 120; index++ {
		go fGet(tmp2)
	}

	for index := 1; index <= 120; index++ {
		<-ch
	}
}
