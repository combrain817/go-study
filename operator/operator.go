package main

import "fmt"

func ex1() {
	var x int32 = 7
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x % y = ", x%y)

	fmt.Println("s * t = ", s*t)
	fmt.Println("s / t = ", s/t)
}

func ex2() {
	var a int = 10
	var b int = 20

	a, b = b, a
	fmt.Println(a, b)
}

func main() {
	ex2()
}
