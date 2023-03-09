package main

import "fmt"

func scanEx() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

func scanfEx() {
	var a int
	var b int

	n, err := fmt.Scanf("%d %d\n", &a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

func scanLn() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

func ex3() {
	var a = 123
	var b int = 4567
	f := 3.14159269

	fmt.Printf("%6d\n", a)
	fmt.Printf("%6d\n", b)
	fmt.Printf("%6.2f\n", f)

}

func main() {
	ex3()
}
