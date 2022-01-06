package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func hoge() {
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}
