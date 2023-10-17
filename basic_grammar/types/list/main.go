// golangの配列の特徴は、要素数を変更できないこと
// 要素数を変更したい場合は、スライスを使用する
package main

import "fmt"

func main(){
	// 型のみ定義してみる
	var list1 [3]int
	fmt.Printf("%T\n", list1) // 型は[3]int, 配列の長さが変わると型も変更することになるので、配列は不変長

	// 明示的に値を定義する
	// 要素数が足りないと、その型における初期値が入る
	var list2 [3]int = [3]int{72, 73}
	fmt.Println(list2) // [72 73 0]

	// 暗黙的に値を定義
	list3 := [3]string{"A", "B", "C"}
	fmt.Println(list3)

	// 要素数を省略して定義する
	list4 := [...]string{"A", "B", "C"}
	fmt.Printf("%T\n", list4)

	// 要素を更新する
	list3[2] = "I"
	fmt.Println(list3)

	// 配列の要素数は同じでも、型が違えば代入することはできない
	// list3 = list2

	// 配列の要素数を調べるlen関数
	fmt.Println(len(list1))
}