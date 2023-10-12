// インターフェイス

package main

import "fmt"

func main(){
	var x interface {}
	fmt.Println(x) //初期値としてnilという特殊な値を持つ

	// {}のため、なんでも代入できるが、演算の対象としては利用できない
	x = "A"
	fmt.Println(x)
	x = [3]int{1,2,3}
	fmt.Println(x)
}