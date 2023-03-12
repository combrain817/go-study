package main

import "fmt"

const Pig int = 0
const Cow int = 1
const Chicken int = 2

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...")
	}
}

func iotaExam() {
	const (
		Red   int = iota
		Blue  int = iota
		Green int = iota
	)

	fmt.Println(Red, Blue, Green)

	const (
		C1 uint = iota + 1
		C2
		C3
	)

	fmt.Println(C1, C2, C3)

}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)

	iotaExam()
}
