package lock

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	l     = sync.Mutex{}
	m     = make(map[int]int)
	rwL   = sync.RWMutex{}
)

//全局变量并发写 导致计数错误
func vari() {
	for i := 0; i < 10000; i++ {
		go func(i int) {
			//defer l.Unlock()
			//l.Lock()
			count++
		}(i)
	}
	fmt.Println(count)
}

//map 并发写 不加锁 fatal error: concurrent map writes
func mp() {
	for i := 0; i < 1000; i++ {
		go func() {
			defer l.Unlock()
			l.Lock()
			m[0] = 0
		}()
	}
}

//不加锁的话 有可能是读的错误的值
func read() {
	defer rwL.RUnlock()
	rwL.RLock()
	fmt.Println("read ", m[0])
}

//如果不加锁 会报错 fatal error: concurrent map writes
func write() {
	defer rwL.Unlock()
	rwL.Lock()
	m[0] = m[0] + 1
}
func rwLock() {
	for i := 0; i < 10000; i++ {
		go read()
	}
	for i := 0; i < 10000; i++ {
		go write()
	}
}

func main() {
	//vari()
	//mp()
	rwLock()
	time.Sleep(3 * time.Second)
}
