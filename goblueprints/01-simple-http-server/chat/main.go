package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// http.Handlerインターフェイスを実装する
type templateHandler struct {
	filename string
	once     sync.Once
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("Fail to start server:", err)
	}
}
