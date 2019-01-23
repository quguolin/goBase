package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

const (
	file1     = "/private/tmp/1.txt"
	file2     = "/private/tmp/2.txt"
	filediff  = "/private/tmp/diff.txt"
	filepatch = "/private/tmp/patch.txt"
)

func bsdiff() (err error) {
	cmd := exec.Command("bsdiff", file1, file2, filediff)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	return cmd.Run()
}

func bspatch() (err error) {
	cmd := exec.Command("bspatch", file2, filepatch, filediff)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	return cmd.Run()
}

func md5Cal(file string) (b []byte) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return h.Sum(nil)
}

func md5File() (result bool) {
	md5File2 := md5Cal(file2)
	md5FilePatch := md5Cal(filepatch)
	return bytes.Equal(md5File2, md5FilePatch)
}

func main() {
	err := bsdiff()
	if err != nil {
		panic(err)
	}
	err = bspatch()
	if err != nil {
		panic(err)
	}
	if ok := md5File(); ok {
		fmt.Println("success!")
	} else {
		fmt.Println("fail!")
	}
}
