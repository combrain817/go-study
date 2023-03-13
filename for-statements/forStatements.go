package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func ex1() {
	i := 1
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
}

func ex2() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요.")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요.")

			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다 \n", number)
		if number%2 == 0 {
			break
		}
	}
	fmt.Println("for문이 종료됬습니다.")
}

func ex3() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	//ex1()
	//ex2()
	ex3()

}
