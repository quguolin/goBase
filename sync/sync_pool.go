package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

var (
	pool *sync.Pool = &sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
)

func BytesBuffer() {
	buf := &bytes.Buffer{}
	buf.WriteString("hello")
	buf.WriteString(fmt.Sprintf("%s", "world"))
	buf.WriteString(time.Now().Format("2016-01-02"))
	buf.WriteTo(ioutil.Discard)
}

func BytesBufferPool() {
	buf := pool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		pool.Put(buf)
	}()
	buf.WriteString("hello")
	buf.WriteString(fmt.Sprintf("%s", "world"))
	buf.WriteString(time.Now().Format("2016-01-02"))
	buf.WriteTo(ioutil.Discard)
}
