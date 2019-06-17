package main

import "testing"

func BenchmarkBytesBuffer(b *testing.B) {
	for i:=0;i<b.N;i++{
		BytesBuffer()
	}
}

func BenchmarkBytesBufferPool(b *testing.B) {
	for i:=0;i<b.N;i++{
		BytesBufferPool()
	}
}
