package main

import (
	"io"
	"testing"
)

var (
	w   io.Writer
	err error
)

func init()  {
	cf := &Config{
		Dir:"/data/log/test/info.log",
	}
	if w, err = NewLog(cf); err != nil {
		panic(err)
	}
}

func TestFLog_Write(t *testing.T) {
	_,err := w.Write([]byte("hello"))
	if err != nil{
		t.Error(err)
	}
}

func BenchmarkFLog_Write(b *testing.B) {
	for i:=0;i<b.N;i++{
		_,err = w.Write([]byte([]byte("hello world")))
		if err != nil{
			b.Error(err)
		}
		_,err = w.Write([]byte("\n"))
		if err != nil{
			b.Error(err)
		}
	}
}
