package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// httpget()
	httppost()
}

func httpget() {
	// parameter
	key1 := "key1"
	param1 := "value1"
	key2 := "key2"
	param2 := "value2"

	// request
	resp, err := http.Get(fmt.Sprintf("http://localhost:12345/index.html?%s=%s&%s=%s", key1, param1, key2, param2))
	if err != nil {
		log.Fatal(err.Error())
	}

	// response
	defer resp.Body.Close()
	buf := make([]byte, 1024)
	i, rerr := resp.Body.Read(buf)
	if err != nil {
		log.Fatal(rerr.Error())
	}
	fmt.Printf("response:%s", string(buf[:i]))
}

func httppost() {
	// parameter
	key1 := "key1"
	param1 := "value1"
	key2 := "key2"
	param2 := "value2"

	// request
	v := url.Values{key1: {param1}, key2: {param2}}
	resp, err := http.PostForm(fmt.Sprintf("http://localhost:12345/upload"), v)
	if err != nil {
		log.Fatal(err.Error())
	}

	// response
	defer resp.Body.Close()
	buf := make([]byte, 1024)
	i, rerr := resp.Body.Read(buf)
	if err != nil {
		log.Fatal(rerr.Error())
	}
	fmt.Printf("response:%s", string(buf[:i]))
}
