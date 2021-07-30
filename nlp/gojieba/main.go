package main

import (
	"fmt"
	"strings"

	"github.com/yanyiwu/gojieba"
)

func main() {
	var s string
	var words []string
	x := gojieba.NewJieba()
	defer x.Free()

	//s = "大白水150ml+乳120g"
	s = "甜莓味20g"
	use_hmm := true
	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("精确模式:", strings.Join(words, "/"))

	//s = "长春市长春药店"
	words = x.Tag(s)
	fmt.Println(words)
	fmt.Println("词性标注:", strings.Join(words, ","))

	//s = "大白水150ml"
	words = x.CutAll(s)
	fmt.Println(s)
	fmt.Println("全模式:", strings.Join(words, "/"))
}
