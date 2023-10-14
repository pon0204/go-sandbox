package main

import "fmt"

// 基本的な関数定義
func add(x int, y int) int {
	return x + y
}

// 型が同じ場合は省略可能
func add2(x, y int) int {
	return x + y
}

// 複数のもどり値を返すことができる
func swap(x, y string) (string, string) {
	return y, x
}

// 返り値に名前をつけることができる
func split(sum int) (x, y  int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func FunctionPractice() {
	a,b := swap("hello", "world")
	fmt.Println(a,b)
}