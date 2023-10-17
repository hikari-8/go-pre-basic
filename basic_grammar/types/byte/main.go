package main

import "fmt"

func main(){
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// stringに変換
	fmt.Println(string(byteA))

	// 文字列をbyteのスライスに変換する
	string1 := []byte("HI")
	fmt.Println(string1)
}