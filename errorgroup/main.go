package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang/sync/errgroup"
)

func httpGet(ctx context.Context, url string) (value string, err error) {
	var (
		req  *http.Request
		resp *http.Response
	)
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return
	}
	req = req.WithContext(ctx)
	client := http.Client{}
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return string(data), nil
}

func main() {
	var (
		result1, result2 string
		err              error
	)
	d := time.Now().Add(time.Duration(1) * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		result1, err = httpGet(ctx, "http://www.google.com/")
		//manually cancel err
		if err != nil {
			log.Println(err)
		}
		return nil
	})
	group.Go(func() error {
		result2, err = httpGet(ctx, "http://www.bilibili.com/")
		return err
	})
	//wait all group to complete
	if err := group.Wait(); err != nil {
		log.Println(err)
		return
	}
	fmt.Println(result1)
	fmt.Println(result2)
}
