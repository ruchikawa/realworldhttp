package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var image []byte

func init() {
	var err error
	image, err = ioutil.ReadFile("./unnamed.jpg")
	if err != nil {
		panic(err)
	}
}

func handleHtml(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher)
	if ok {
		pusher.Push("/image", nil)
	}

	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><body><img src="/image"></img></body></html>`)
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpg")
	w.Write(image)
}

func main() {
	http.HandleFunc("/", handleHtml)
	http.HandleFunc("/image", handleImage)
	fmt.Println("start http listening :18443")
	err := http.ListenAndServeTLS(":18443", "server.crt", "server.key", nil)
	fmt.Println(err)
}
