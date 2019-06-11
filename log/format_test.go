package main

import (
	"os"
	"testing"
)

var (
	f = NewFormat()
)

func TestFormat_Format(t *testing.T) {
	vs := make(map[string]interface{}, 0)
	vs["Info"] = "test"
	f.Format(os.Stdout, vs)
}
