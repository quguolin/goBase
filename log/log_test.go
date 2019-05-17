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
	if w, err = NewLog(); err != nil {
		panic(err)
	}
}

func BenchmarkFLog_Write(b *testing.B) {
	for i:=0;i<b.N;i++{
		w.Write([]byte([]byte("hello world")))
		w.Write([]byte("\n"))
	}
}
