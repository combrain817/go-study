package main

import "fmt"

// 변수명 지정해 반환하기
func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}

	result = a / b
	success = true
	return
}

func Multiple(a int, b int) int {
	return a * b
}

func F(n int) int {

	if n < 2 {
		return n
	}

	return F(n-2) + F(n-1)
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)

	d, success := Divide(9, 0)
	fmt.Println(d, success)

	result := Multiple(3, 5)
	fmt.Println(result)

	F(9)
}
