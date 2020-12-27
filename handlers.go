package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)
var shallowCopy map[interface{}]interface{}

func init() {
	shallowCopy = make(map[interface{}]interface{})
}

func copyKeyValue(key interface{}, value interface{}) bool {
	shallowCopy[key] = value
	return true
}

func listKeys(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	dict.Range(copyKeyValue)
	go fmt.Println("Listing Keys")
	fmt.Fprint(w, shallowCopy)
}

func deleteKeyValue(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	dict.Delete(key)
	go fmt.Println("Deleting Key " + key)
	fmt.Fprint(w)
}

func getKeyValue(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	value, _ := dict.Load(key)
	go fmt.Println("Getting Key " + key + " w/ value ")
	fmt.Fprint(w, value)
}

func setKeyValue(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	value := ps.ByName("value")
	go fmt.Println("Setting Key " + key + " to " + value)
	dict.Store(key, value)
	fmt.Fprint(w, value)
}
