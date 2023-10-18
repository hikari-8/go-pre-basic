package main

// sivchariさん
// スケッチの輪郭を書くことを意識しよう
// 自分が本当に興味がある周辺を読んで行って、1つ1つ頭の中の付箋を作っていくイメージ
// 最初にコードを書き始める時はできるだけハードルを下げること！

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
