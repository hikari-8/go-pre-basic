package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	//乱数の種を初期化
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		randomNum := randomNumMaker()
		res := omikuji(randomNum)
		fmt.Fprint(w, "おみくじの結果は...", res)
	})
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("Fail to start:", err)
	}
}

func omikuji(randomNum int) string {
	var res string
	switch randomNum {
	case 0:
		res = "大吉です"
	case 1:
		res = "中吉です"
	case 2:
		res = "吉です"
	case 3:
		res = "凶です"
	}
	return res
}

func randomNumMaker() int {
	res := rand.Intn(4)
	return res
}

//条件
//Webアプリ化してみよう
//HTTPサーバを作成する
//リクエストが来たらおみくじの結果を返す
//乱数の種は1回だけ初期化する
//HTTPサーバを起動する前に初期化する
