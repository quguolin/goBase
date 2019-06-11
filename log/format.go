package main

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

const formatP = "%T %L"

func timeFormat(map[string]interface{}) string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func logFormat(key string) func(map[string]interface{}) string {
	return func(vs map[string]interface{}) string {
		if v, ok := vs[key]; ok {
			if s, ok := v.(string); ok {
				return s
			}
			return fmt.Sprint(v)
		}
		return ""
	}
}

//t time
//l log
var formatMap = map[string]func(map[string]interface{}) string{
	"T": timeFormat,
	"L": logFormat("Info"),
}

type format struct {
	funcs []func(map[string]interface{}) string
}

func NewFormat() *format {
	var (
		formatStr string = formatP
	)
	fs := &format{}
	for i := 0; i < len(formatStr); i++ {
		if formatStr[i] != '%' {
			continue
		}
		f, ok := formatMap[string(formatStr[i+1])]
		if ok {
			fs.funcs = append(fs.funcs, f)
		}
	}
	return fs
}

func (format *format) Format(w io.Writer, vs map[string]interface{}) error {
	buf := &bytes.Buffer{}
	for _, f := range format.funcs {
		buf.WriteString(f(vs))
	}
	_, err := buf.WriteTo(w)
	return err
}
