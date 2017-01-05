package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//	"reflect"
	"time"
)

type ticket struct {
	time   time.Time
	during uint
}

type wxError struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type wxToken struct {
	AccessToken string `json:"access_token"`
	Expires     int    `json:"expires_in"`
}
type wxAccessToken struct {
	wxError
	wxToken
}

var f *os.File
var err1 error

func main() {
	//	_ := time.Now()
	var tmp wxToken

	//	value := reflect.TypeOf(time)
	//	fmt.Println(value)
	if token := getAccessToken(); token.Errmsg != "" {
		fmt.Println(token.Errmsg)
	} else {
		tmp.AccessToken = token.AccessToken
		tmp.Expires = token.Expires

		value, err := json.Marshal(tmp)
		if err != nil {
			log.Fatal(err)
		}

		filename := "token.txt"

		n := writeFile(filename, value)
		fmt.Printf("写入%d个字符", n)

	}
}

func getAccessToken() wxAccessToken {
	appid := ""
	secret := ""
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	value, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	//	fmt.Printf("%s", value)
	var token wxAccessToken
	err = json.Unmarshal(value, &token)
	if err != nil {
		log.Fatal(err)
	}

	return token
}

func writeFile(filename string, content []byte) int {
	if checkFileExist(filename) {
		//文件存在
		f, err1 = os.OpenFile(filename, os.O_WRONLY, 0666)
		if err1 != nil {
			log.Fatal(err1)
		}
	} else {
		//文件不存在
		f, err1 = os.Create(filename)
		if err1 != nil {
			log.Fatal(err1)
		}
	}

	n, err := io.WriteString(f, string(content))
	if err != nil {
		log.Fatal(err)
	}

	return n
}
func checkFileExist(filename string) bool {
	exist := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}

	return exist

}
