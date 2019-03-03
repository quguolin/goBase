package main

import (
	"fmt"
	"go-common/app/service/ops/log-agent/pkg/bufio"
)

func interType(arg interface{}) {
	switch arg := arg.(type) {
	case string:
		fmt.Println("string")
	case []byte:
		fmt.Println("byte")
	case int:
		fmt.Println("int")
	case int64:
		fmt.Println("int64")
	case float64:
		fmt.Println("float64")
	case bool:
		if arg {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	case nil:
		fmt.Println("nil")

	}
}

func main() {
	var (
		a string  = "string"
		b []byte  = []byte("byte")
		c int     = 10
		d int64   = 10
		e float64 = 10
		f bool    = true
	)
	args := []interface{}{
		a, b, c, d, e, f, nil,
	}
	for _, v := range args {
		interType(v)
	}

	w := bufio.NewWriter(nil)
	w.Write([]byte("a"))
	w.WriteString("b")
}
