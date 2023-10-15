package main

import "fmt"

// 初期値を与えずに宣言すると、ゼロ値が与えられる
func Zero() {
	// 0
	var i int 
	// 0
	var f float64
	// false
	var b bool
	// ""
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}