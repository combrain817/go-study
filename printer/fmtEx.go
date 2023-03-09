package main

import "fmt"

func printEx() {
	var a int = 10
	var b int = 20
	var f float64 = 3299438743.8297

	fmt.Print("a: ", a, "b :", b)
	fmt.Println("a: ", a, "b: ", b, "f: ", f)
	fmt.Printf("a: %d b: %d f:%f \n", a, b, f)
}

func printFormatEx() {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%5d, %5d\n", a, b)
	fmt.Printf("%05d %05d\n", a, b)
	fmt.Printf("%-5d, %-05d\n", a, b)

	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)

	var d = 324.13455
	var e = 3.14

	fmt.Printf("%08.2f\n", d)
	fmt.Printf("%08.2g\n", d)
	fmt.Printf("%8.5g\n", d)
	fmt.Printf("%f\n", e)

}
