package main

import (
	"fmt"
)

func GetSquare(num int, ch chan<- int) {
	ch <- num * num
}

func main() {
	arr := [...]int{2, 4, 6, 8, 10}

	ch := make(chan int, len(arr))

	for _, i := range arr {
		go GetSquare(i, ch)
	}

	sum := 0
	for range arr {
		sum += <-ch
	}
	close(ch)

	fmt.Println(sum)
}
