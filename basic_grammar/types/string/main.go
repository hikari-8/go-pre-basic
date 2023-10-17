// string
package main

import "fmt"

func main(){
	// 複数行にわたって出力するにはバッククオーテーションを使用する
	fmt.Println(`
	test
	test
	test
	`)
	// "を文字列としてコンソールに出力する
	fmt.Println("\"") // or
	fmt.Println(`"`)

// 文字列型はバイト配列の集まり＝アスキーコードで表現されている
var i1 string = "Hello Gopher!"
fmt.Println(i1[0]) //byteが出力される
fmt.Println(string(i1[0]))  // stringに変換

}