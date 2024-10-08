package main

import (
	"fmt"
	"sync"
)

func writeArray(arr []int, ch chan<- int, wg *sync.WaitGroup) {
	for _, i := range arr {
		ch <- i
	}
	close(ch)

	wg.Done()
}

func square(chFrom <-chan int, chTo chan<- int, wg *sync.WaitGroup) {
	for {
		val, ok := <-chFrom
		if !ok {
			break
		}
		chTo <- val * val
	}
	close(chTo)

	wg.Done()
}

func printStdout(ch <-chan int, wg *sync.WaitGroup) {
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(val)
	}

	wg.Done()
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(1)
	go writeArray(arr, ch1, &wg)

	wg.Add(1)
	go square(ch1, ch2, &wg)

	wg.Add(1)
	go printStdout(ch2, &wg)

	wg.Wait()
}
