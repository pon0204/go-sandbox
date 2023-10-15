package main

import "fmt"

const Pi = 3.14

func Constants() {
	// 定数は:=を使って宣言できない。 文字(character)、文字列(string)、boolean、数値(numeric)のみが定数で使える。
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
