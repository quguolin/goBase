package limit

import (
	"fmt"
	"testing"
	"time"
)

func TestLimit_Reset(t *testing.T) {
	duration, err := time.ParseDuration("4s")
	if err != nil {
		panic(err)
	}
	l := New(2, duration)
	second, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	for t := range time.Tick(second) {
		fmt.Println(t.Unix())
		if l.isLimit() {
			fmt.Println("ok")
		} else {
			fmt.Println("refuse")
		}
	}
}
