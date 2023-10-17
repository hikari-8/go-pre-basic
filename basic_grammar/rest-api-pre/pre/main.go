package main

import (
	"log"
	"net/http"
)

func main(){

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		body := []byte(`<html><head><title>Go pre</title></head><body><p>This is Home Page</p></body></html>`)
		w.Write(body)
	})

	http.HandleFunc("/sub", func(w http.ResponseWriter, r *http.Request){
		subbody := []byte(`<html><head><title>Go pre</title></head><body><p>This is Sub Page</p></body></html>`)
		w.Write(subbody)
	})

	if err := http.ListenAndServe(":8000", nil); err!= nil {
		log.Fatal("fail", err)
	}

}