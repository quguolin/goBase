package benchMark

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"sync"
	"testing"
)

type TestingData struct {
	Data string `json:"data"`
	Key  string `json:"key"`
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}

func BenchmarkWithPool(b *testing.B) {
	data := TestingData{
		Data: "data",
		Key:  "key",
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buf := GetBuffer()
		_ = json.NewEncoder(buf).Encode(&data)
		io.Copy(ioutil.Discard, buf)
		PutBuffer(buf)
	}
}

func BenchmarkWithoutPool(b *testing.B) {
	data := TestingData{
		Data: "data",
		Key:  "key",
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buf := &bytes.Buffer{}
		_ = json.NewEncoder(buf).Encode(&data)
		io.Copy(ioutil.Discard, buf)
	}
}
