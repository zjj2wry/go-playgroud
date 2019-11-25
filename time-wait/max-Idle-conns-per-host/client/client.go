package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 50; i++ {
		go request()
	}
	time.Sleep(100 * time.Second)
}

func request() {
	Customize the Transport to have larger connection pool
	defaultRoundTripper := http.DefaultTransport
	defaultTransportPointer, ok := defaultRoundTripper.(*http.Transport)
	if !ok {
		panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
	}
	defaultTransport := defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 50

	client := &http.Client{Transport: defaultTransport}
	// client := &http.Client{}

	for {
		var jsonStr = []byte(`{"a":"a"}`)
		req, err := http.NewRequest("POST", "http://192.168.50.4:7070", bytes.NewBuffer(jsonStr))
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("response Body:", string(body))
		resp.Body.Close()
	}
}
