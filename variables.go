package main

import "fmt"

var c, python, java bool

var i, j int = 1,2

func Variables() {
	var i, j int = 1,2
	// 暗黙的な型宣言。varを省略出来る。関数内でのみ使用可能。
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
	
}