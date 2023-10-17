package main

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