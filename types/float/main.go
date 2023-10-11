// 不動小数点型

package main

import "fmt"

func main(){
	var f1 float64 = 8.8
	// 暗黙的に定義する場合は、float64として推論される 
	// intと違って、環境依存しない
	f2 := 2.1
	fmt.Printf("%T, %T\n", f1, f2)
	fmt.Println(f1+f2)

	//float32もあるが、型が違うとコンパイルエラーに怒られる
	// float32はあまり使わない
	var f3 float32 = 4.5
	fmt.Printf("%T, %T\n", f1, f3)
	// fmt.Println(f1+f3)

	//型変換も可能
	fmt.Printf("%T\n", float64(f3))

	// 正の無限大と負の無限大,  数値として表現できない結果
	// エラーや特殊な状況を示すために使用される
	zero := 0.0
	pinf :=  1.1 / zero // +Inf

	nanf := -1.1 / zero // -Inf

	nan := 0.0 / zero //NaN

	fmt.Println(pinf, nanf, nan) 
}