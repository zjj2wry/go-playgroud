package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go log.Println(http.ListenAndServe("localhost:6060", nil))
}
