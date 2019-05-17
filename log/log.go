package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sync"
)

type fLog struct {
	fp     *os.File
	stdLog *log.Logger
	ch     chan *bytes.Buffer
	wg     *sync.WaitGroup
}

func NewLog() (*fLog, error) {
	fp, err := os.OpenFile("/data/log/test/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	stdLog := log.New(os.Stderr, "fileLog", log.LstdFlags)
	ch := make(chan *bytes.Buffer, 10)
	wg := &sync.WaitGroup{}
	f := &fLog{
		fp:     fp,
		stdLog: stdLog,
		ch:     ch,
		wg:     wg,
	}
	f.wg.Add(1)
	go f.daemon()
	return f, nil
}

func (l *fLog) close() error {
	close(l.ch)
	l.wg.Wait()
	return nil
}

func (l *fLog) daemon() {
	defer l.wg.Done()
	for {
		select {
		case v, ok := <-l.ch:
			if ok {
				l.write(v.Bytes())
			}
		}
	}
}

func (l *fLog) Write(p []byte) (int, error) {
	buf := &bytes.Buffer{}
	buf.Write(p)
	select {
	case l.ch <- buf:
		return len(p), nil
	default:
		return 0, fmt.Errorf("log channel is full")
	}
}

func (l *fLog) write(p []byte) error {
	_, err := l.fp.Write(p)
	if err != nil {
		l.stdLog.Printf("write log error:%v", err)
		return err
	}
	return nil
}

//func main() {
//	var (
//		w   io.Writer
//		err error
//	)
//	if w, err = NewLog(); err != nil {
//		panic(err)
//	}
//	for {
//		w.Write([]byte("test"))
//		w.Write([]byte("\n"))
//		time.Sleep(time.Second)
//	}
//}
