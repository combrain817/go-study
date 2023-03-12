package main

import "fmt"

func ex1(num int) {

	switch num {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")

	}
}

func ex2(day string) {

	switch day {
	case "monday", "tuesday":
		fmt.Println("월, 화요일은 수업 가는 날입니다.")
	case "wednesday", "thursday", "friday":
		fmt.Println("수, 목, 금요일은 실습 가는 날입니다.")
	}
}

func ex3(temp int) {

	switch true {
	case temp < 10, temp > 30:
		fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하세요")
	case temp >= 15 && temp < 25:
		fmt.Println("야외 활동하기 좋은 날씨입니다")
	default:
		fmt.Println("따뜻합니다")
	}
}

func getMyAge() int {
	return 29
}

func ex4() {
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is ", age)
	}
}

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func ex5(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func ex6(num int) {
	switch num {
	case 1:
		fmt.Println("num == 1")
		break
	case 2:
		fmt.Println("num == 2")
	case 3:
		fmt.Println("num == 3")
		fallthrough
	case 4:
		fmt.Println("num == 4")
	case 5:
		fmt.Println("num == 5")
	default:
		fmt.Println("num > 5")
	}
}

func main() {
	ex1(3)
	ex2("friday")
	ex3(25)
	ex4()
	fmt.Println(ex5(Green))
	ex6(3)
}
