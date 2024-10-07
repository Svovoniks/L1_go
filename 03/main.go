package main

import (
	"fmt"
)

func GetSquare(num int, ch chan int) {
	ch <- num * num
}

func main() {
	arr := [...]int{2, 4, 6, 8, 10}

	var chArr []chan int

	for _, i := range arr {
		ch := make(chan int)
		chArr = append(chArr, ch)

		go GetSquare(i, ch)
	}

	sum := 0

	for _, i := range chArr {
		sum += <-i
	}

    fmt.Println(sum)
}
