package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func ex1() {
	var house House
	house.Address = "서울시"
	house.Size = 28
	house.Price = 9.8
	house.Type = "아파트"

	fmt.Println("주소:", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격: %.2f억 원\n", house.Price)
	fmt.Println("타입: ", house.Type)

}

type Student struct {
	Age   int
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("나이: %d 번호: %d 점수:%.2f\n", s.Age, s.No, s.Score)
}

func ex2() {
	var student = Student{15, 23, 88.2}

	student2 := student

	PrintStudent(student)
	PrintStudent(student2)

	student.Age = 100

	PrintStudent(student)
	PrintStudent(student2)

}

func main() {
	ex2()
}
