package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

//graphviz
//内存 --inuse_objects 每个函数分配的对象数 --alloc-space 查看分配的内存空间大小  --nodefration=0.05 子函数使用的 CPU、memory 不超过 5%，就忽略它，不要显示在图片中
//go tool pprof http://localhost:6060/debug/pprof/heap --inuse_objects  --alloc-space --nodefration=0.05
//go tool pprof http://localhost:6060/debug/pprof/profile
func counter() {
	list := []int{1}
	c := 1
	for i := 0; i < 10000000; i++ {
		httpGet()
		c = i + 1 + 2 + 3 + 4 - 5
		list = append(list, c)
	}
	fmt.Println(c)
	fmt.Println(list[0])
}

func work(wg *sync.WaitGroup) {
	for {
		counter()
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func httpGet() int {
	queue := []string{"start..."}
	resp, err := http.Get("http://www.163.com")
	if err != nil {
		// handle error
	}

	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	queue = append(queue, string(body))
	return len(queue)
}

func main() {
	flag.Parse()
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 100; i++ {
		go work(&wg)
	}

	wg.Wait()
	time.Sleep(3 * time.Second)
}
