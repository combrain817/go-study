package main

import "fmt"

func ex1() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}
}

func ex2() {
	nums := [...]int{10, 20, 30, 40, 50}

	nums[2] = 300

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}

func ex3() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t {
		fmt.Println(i, v)
	}
}

func ex4() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a

	fmt.Println()

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
	fmt.Println()
	a[1] = 100

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

}

func main() {
	ex4()
}
