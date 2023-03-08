package main

import "fmt"

func variableEx() {
	var a int = 3
	var b int
	var c = 4
	d := 5

	fmt.Println(a, b, c, d)
}

func typeEx() {
	a := 3
	var b float64 = 3.5

	var c int = int(b)
	d := float64(a * c)

	var e int64 = 7
	f := int64(d) * e

	var g int = int(b * 3)
	var h int = int(b) * 3

	fmt.Println(g, h, f)
}

func typeEx2() {
	var a int16 = 3456
	var c int8 = int8(a)

	fmt.Println(a)
	fmt.Println(c)
}

var g int = 10

func rangeEx() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

}

func main() {
	rangeEx()
}
