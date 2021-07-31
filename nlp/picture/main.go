package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("/Users/quguoiln/Downloads/43a76e2d-cd4d-4b83-98e6-8ea67b954a1b.jpg.a.jpeg")
	text, _ := client.Text()
	fmt.Println(text)
}
