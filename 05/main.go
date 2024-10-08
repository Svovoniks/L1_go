package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Cousume(channel <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case i := <-channel:
			fmt.Println("received", i)
		case <-ctx.Done():
			fmt.Println("exited consumer")
			return
		}
	}
}

func Produce(channel chan<- int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	counter := 0
	for {
		select {
		case channel <- counter:
			fmt.Println("sent", counter)
		case <-ctx.Done():
			fmt.Println("exited producer")
			return
		}

		counter++
	}
}

func main() {
	var exitAfter int
	if _, err := fmt.Scan(&exitAfter); err != nil || exitAfter < 1 {
		fmt.Println("Not a valid number of seconds")
		return
	}

	channel := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(exitAfter)*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(2)
	go Produce(channel, ctx, &wg)
	go Cousume(channel, ctx, &wg)

	<-ctx.Done()

	wg.Wait()
}
