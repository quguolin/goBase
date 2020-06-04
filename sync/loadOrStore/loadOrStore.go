package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	m     = sync.Map{}
	count int64
	tmp   int64
)

func Add(supplierid string) {
	_, ok := m.LoadOrStore(supplierid, 1)
	if ok {
		fmt.Println("有更新 跳过")
		atomic.AddInt64(&tmp, 1)
		return
	}
	// todo 业务逻辑
	count++
	m.Delete(supplierid)
}

func main() {
	var wg sync.WaitGroup
	var supplierid = "10"
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Add(supplierid)
		}()
	}
	wg.Wait()
	fmt.Println(count)
	fmt.Println(tmp)
}
