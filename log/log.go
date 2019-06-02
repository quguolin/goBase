package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
)

const flagClosed = 1

type fLog struct {
	fp     *os.File
	stdLog *log.Logger
	ch     chan *bytes.Buffer
	wg     *sync.WaitGroup
	closed int32
}

type Config struct {
	Dir string
}

func NewLog(cf *Config) (*fLog, error) {
	dir := filepath.Dir(cf.Dir)
	name := filepath.Base(cf.Dir)
	if name == ""{
		return nil,fmt.Errorf("name can not empty")
	}
	fInfo,err := os.Stat(dir)
	if err == nil && !fInfo.IsDir(){
		return nil,fmt.Errorf("%s existed and not a directory ",dir)
	}
	if os.IsNotExist(err){
		if err = os.MkdirAll(cf.Dir,0755);err != nil{
			return nil,fmt.Errorf("create file err %s",err)
		}
	}
	fp, err := os.OpenFile(cf.Dir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	stdLog := log.New(os.Stderr, "fileLog", log.LstdFlags)
	ch := make(chan *bytes.Buffer, 1024*100)
	wg := &sync.WaitGroup{}
	f := &fLog{
		fp:     fp,
		stdLog: stdLog,
		ch:     ch,
		wg:     wg,
	}
	f.wg.Add(1)
	go f.consume()
	return f, nil
}

func (l *fLog) consume() {
	defer l.wg.Done()
	bf := &bytes.Buffer{}
	tk := time.NewTicker(10*time.Millisecond)
	for {
		select {
		case v, ok := <-l.ch:
			if ok {
				bf.Write(v.Bytes())
			}
		case <-tk.C:
			if bf.Len()>0{
				l.write(bf.Bytes())
				bf.Reset()
			}
		}
		if atomic.LoadInt32(&l.closed) != flagClosed {
			continue
		}
		//log closed and read all from buffer and channel
		l.write(bf.Bytes())
		for bf := range l.ch{
			l.write(bf.Bytes())
		}
		return
	}
}

func (l *fLog) Write(p []byte) (int, error) {
	//buf := &bytes.Buffer{}
	//log closed
	if atomic.LoadInt32(&l.closed) == flagClosed {
		l.stdLog.Printf("%s",p)
		return 0,fmt.Errorf("log is closed")
	}
	buf := bytes.NewBuffer(make([]byte, 0, len(p)))
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

func (l *fLog) Close() error {
	atomic.StoreInt32(&l.closed, flagClosed)
	close(l.ch)
	l.wg.Wait()
	return nil
}
