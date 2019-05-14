package main

import (
	"bytes"
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
	w.Write([]byte(bytes.Repeat([]byte("testtesttesttesttesttest"),1000000)))
	w.Write([]byte("\n"))
}
