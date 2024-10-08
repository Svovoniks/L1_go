package main

import (
	"fmt"
	"sync"
)

func PrintSquare(num int, wg *sync.WaitGroup) {
	fmt.Println(num * num)
	wg.Done()
}

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, i := range arr {
		wg.Add(1)
		go PrintSquare(i, &wg)
	}

	wg.Wait()
}
