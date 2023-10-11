// int
package main

import "fmt"

func main (){
	// intに何も指定しない場合は、intSizeはアーキテクチャに応じて32か64のどちらかになる
  var i int = 100 
	fmt.Printf("%T\n", i) //%Tは書式指定詞

	// メモリの最適化や明示的に型を指定するために、int8, int16, int32, int64など複数の整数型が提供されている 
	var i2 int64 = 200
	fmt.Println(i2)

	// 別の整数型に代入したり計算しようとすると怒られる
	// i2 = i
	// fmt.Println(i + i2) 

	//このような場合は ${int型}(変数)で変換できる
	fmt.Println(i + int(i2))

}