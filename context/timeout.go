package main

import (
	"context"
	"net/http"
	"time"
)

func httpGet(ctx context.Context) (err error) {
	var (
		req *http.Request
	)
	req, err = http.NewRequest("GET", "https://www.google.com.hk/", nil)
	if err != nil {
		return
	}
	req = req.WithContext(ctx)
	client := http.Client{}
	if _, err = client.Do(req); err != nil {
		return
	}
	return
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := httpGet(ctx); err != nil {
		panic(err)
	}
}
