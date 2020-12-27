package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"sync"
)

var dict sync.Map
var portNumber string

func main() {
	// Get arguments
	args := os.Args[1:]
	if len(args) > 0 {
		portNumber = args[0]
	} else {
		portNumber = "8080"
	}

	// test
	dict.Store("Bell_Labs", "40.68433, -74.39967")
	fmt.Println(dict.Load("Bell_Labs"))


	router := httprouter.New()
	router.GET("/list", listKeys)
	router.GET("/delete/:key", deleteKeyValue)
	router.GET("/get/:key", getKeyValue)
	router.GET("/set/:key/:value", setKeyValue)
	http.ListenAndServe(":" + portNumber, router)
}


