package main

import "fmt"

// 定数は後から値を変更することができない
// 後から値を変更しない場合は定数を使った方がセキュリティー上good

// 基本的に定数はグローバルな値として関数外に定義され、別のpackageからも参照される
// 大文字にすると、グローバルな変数
// 小文字にすると、定義したパッケージ内でしか呼び出せない

const Pi = 3.14

// 一括で定義する, 値を省略して定義する
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// 定義する値の最大値がない
// var Bigint int = 9223372036854775807 +1 int型の最大値を超えてしまうのでoverflowになってしまう
const Bigint = 9223372036854775807 +1

// 整数を0から連番で定義する
const (
	H = iota
	I
	J
)

func main(){
	fmt.Println(Pi)
	fmt.Println(A,B,C,D,E,F)
	fmt.Printf("%T\n", Bigint-1)
	fmt.Println(Bigint -1) //Println関数の引数として渡す時は、intの最大値を超えるとエラーが発生する
	fmt.Println(H, I, J)
}