package main

// slice
// 配列との違いは、要素数を指定しないこと

import "fmt"

func main(){
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{1,2,3,4,5}
	fmt.Println(sl2)

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	// make関数でsliceを生成することもできる
	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl5 := sl2[2]
	fmt.Println(sl5)

	sl6 := sl2[2:3]
	fmt.Println(sl6)

	sl7 := sl2[:2]
	fmt.Println(sl7)

	sl8 := sl2[2:]
	fmt.Println(sl8)

	// 全て
	sl9 := sl2[:]
	fmt.Println(sl9)

	// 最後だけ
	sl10 := sl2[len(sl2)-1]
	fmt.Println(sl10)

	// 最初と最後以外の部分を出力する
	sl11 := sl2[1:len(sl2)-1]
	fmt.Println(sl11)

  //　可変長であることを証明しているappend,make,len,cap
	sl12 := append(sl2, 10, 11, 12)
	fmt.Println(sl12)
	
	sl13 := make([]int, 5, 6)
	fmt.Println(sl13, cap(sl13))

	// 要領以上の要素が追加されるとメモリの消費が倍になってしまいます。
  // メモリーを気にするような開発の場合は、容量にも気をつけます。
	// 良質なパフォーマンスを実現するには、容量の管理も気にします。
	sl13 = append(sl13, 1, 2, 3, 4)
	fmt.Println(sl13, cap(sl13))
	copyfunc()
}

func copyfunc(){
	// 参照型の特製上、このような実装では元の変数も変わってしまう
	s := []int{100, 200}
	s1 := s
	s1[0] = 1000
	fmt.Println(s) 

	// プリミティブ型
	var i int = 5
	i2 := i
	i2 = 10
	fmt.Println(i, i2)

	// よって、copy関数を使う
	sl := []int{1,2,3,4,5}
	sl2 := make([]int, 5,10)
	fmt.Println(sl, sl2)
	n := copy(sl2, sl)
	fmt.Println(n, sl2)

}