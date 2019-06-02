package main

import (
	"context"
	"io"
	"sync"
	"testing"
)

var (
	w   io.Writer
	handle Handler
	err error
	once sync.Once
)

func init()  {
	cf := &Config{
		Dir:"/data/log/test/infoIO.log",
	}
	if w, err = NewLog(cf); err != nil {
		panic(err)
	}
	cf.Dir = "/data/log/test/infoHandle.log"
	if handle,err = NewLog(cf);err != nil{
		panic(err)
	}
}

func BenchmarkFLog_Write(b *testing.B) {
	for i:=0;i<b.N;i++{
		_,err = w.Write([]byte([]byte("info:hello world\n")))
		if err != nil{
			b.Error(err)
		}
	}
}

func BenchmarkFLog_Log(b *testing.B) {
	for i:=0;i<b.N;i++{
		handle.Log(context.Background(),Field{Key:"info",Value:"hello world"})
	}
}
