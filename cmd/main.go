package main

import (
	"goprojects/cms"
	"net/http"
)

func main() {
	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/new", cms.HandleNew)
	http.HandleFunc("/post", cms.ServePost)
	http.HandleFunc("/page", cms.ServePage)
	err := http.ListenAndServe(":3000", nil)
	panic(err)
}
