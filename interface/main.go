package main

// 質問
// ちょっとテキスト復習していたのですが、interfaceの便利さ、使うべき場所の解像度が低いなあと感じたので質問させていただきます！
// まず、現時点でのインターフェイスに対する理解はこんな感じです。
// https://cloudsmith.co.jp/blog/backend/go/2021/08/1847845.html　参考になるかも
// ```
// type Hoge interface {
// 	FuncA()
// 	FuncB()
// }

// type Foo struct {}

// func (f *Foo) FuncA() {}
// func (f *Foo) FuncB() {}

// func main() {
// 	var hoge Hoge
// 	hoge = Foo{..}
// }
// ```
// 1. インターフェイスはメソッドの集合体
// 2. インターフェイスに構造体を代入できる
// 3. 関数の凝集度を高くできるとか、main関数の見栄えがよくなるぐらいしか思いつかなかったです。。。
// 特にインターフェイスに構造体を代入するメリットみたいなものが理解できてないです🤔
// 具体的なユースケースを知れば解像度があがるのかなと思ったので、ぜひ教えていただきたいです！

import "fmt"

// interface
// 具体的にどのように動作するかではなく、どのような動作を持っているべきか（どのメソッドを持つべきか）を示すもの

// インターフェイスの定義
	type Talker interface {
		Talk() string
	}

	// DogもCatも、それぞれTalkメソッドを持っているので、両方ともTalkerインターフェイスを実装しているとみなされる

	type Dog struct {}
	type Cat struct {}

	// インターフェイスの実装
	func (d Dog) Talk() string {
		return "Worf!"
	}

	func (c Cat) Talk() string {
		return "Meow!"
	}

	func MakeSound(t Talker) {
	// インターフェイスを利用した関数を作成する
	// この関数はTalkerインターフェイスを引数として取るため、DogもCatもこの関数に渡すことができる
	fmt.Println(t.Talk())
	}

	func main (){
		d:= Dog{}
		c:=Cat{}
		MakeSound(d)
		MakeSound(c)
	}