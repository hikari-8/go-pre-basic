package main

import "fmt"

func main() {
	// 明示的に変数名と型を指定する方法
	// var 変数名　型 = 値
	var i int = 100
	fmt.Println(i)

	// 2つの変数を同時に定義する
	var t, f bool = true, false
	fmt.Println(t, f)

	// 型の違う変数を一括で定義したい
	var (
		i2 int = 200
		str2 string = "hello!!"
	)

	fmt.Println(i2, str2)

	// 値を代入せずに、変数名を定義する
	// この場合, 値には初期値が入る
	var i3 int
	var s3 string
	var tf bool
	fmt.Println(i3, s3, tf)

	// 値も代入できる
	i3 = 400
	fmt.Println(i3)

	// 暗黙的に変数を定義する
	// varや型定義がいらないが、関数内でしか定義することができない
	i4 := 500
	fmt.Println(i4)
	test()
}

// 明示的な変数定義は関数外でも定義できる
var i5 int = 600
// i5 := 666

func test (){
	fmt.Println(i5)
}

// 明示的or暗黙的に関数指定するのはどちらが推奨されるか
// 1. 関数内か関数外で判断
// 2. 基本的には明示的に型を指定して、可読性を向上させるほうが良い