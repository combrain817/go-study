package main

import "fmt"

func ex1(light string) {

	if light == "green" {
		fmt.Println("길을 건넌다")
	} else {
		fmt.Println("대기한다")
	}
}

func ex2(temp int) {

	if temp > 28 {
		fmt.Println("에어컨을 켠다")
	} else if temp <= 3 {
		fmt.Println("히터를 켠다")
	} else {
		fmt.Println("대기한다")
	}
}

func getMyAge() (int, bool) {
	return 29, true
}

func ex3() {

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	//fmt.Println("Your age is", age)
}

func main() {
	ex1("green")

	ex2(1)

	ex3()
}
