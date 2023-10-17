package main

import "fmt"

func main(){
	m := map[string]int{
		"Apple": 100,
		"Banana": 50,
		"Orange": 	200,
		"Strawberry": 500,
		"Pineapple": 1000,
	}

	res1 := m["Apple"]+m["Strawberry"]
	
	fmt.Println(res1)

	res2:= 0

	for _, val := range m {
		res2 += val*2
	}
	fmt.Println(res2)
}