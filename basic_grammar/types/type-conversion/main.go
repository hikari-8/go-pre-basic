package main

//型変換

import (
	"fmt"
	"strconv"
)


func main(){
	var i int = 1
	f1 := float64(i)
	fmt.Println(f1)
	fmt.Printf("%T\n",f1)

	// floatからintに直すと、小数点以下が切り捨てになる
	f2 := 3.5
	fmt.Println(int(f2))

	// stringからintに変換
	var s2 string = "200"
	integer, _ := strconv.Atoi(s2)
	fmt.Printf("%T\n", integer)

	// intからstringに変換
	// strconv.Atoiメソッドは2つの返り値を返すが、strconv.Itoaは1つの返り値を返す
	var i2 int = 300
	string2 := strconv.Itoa(i2)
	fmt.Printf("%T\n", string2)

	// byte配列から文字列への変換
	s3 := "Hello World!"
	b := []byte(s3)
	fmt.Println(b)

	// 文字列からbyte配列への変換
	fmt.Println(string(b))

}