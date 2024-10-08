package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	chan1 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-chan1:
				fmt.Println(1, "stopped")
				return
			default:
				fmt.Println(1, "working")
			}
			time.Sleep(time.Second)
		}

	}()

	// Завершим рутину записью значения в канал
	chan1 <- 1
	wg.Wait()

	chan2 := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case v, ok := <-chan2:
				if !ok {
					fmt.Println(2, "stopped")
					return
				}
				fmt.Println(v)
			default:
				fmt.Println(2, "working")
			}
			time.Sleep(time.Second)
		}

	}()

	// Завершим рутину закрытием канала
	close(chan2)
	wg.Wait()

	// Завершим рутину по таймеру
	chan3 := time.After(time.Second)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-chan3:
				fmt.Println(3, "stopped")
				return
			default:
				fmt.Println(3, "working")
			}
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()

	// Завершим рутину по таймауту
	ctx4, cancel4 := context.WithTimeout(context.Background(), time.Second)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx4.Done():
				fmt.Println(4, "stopped")
				return
			default:
				fmt.Println(4, "working")
			}
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	cancel4()

	// Завершим рутину с помощью отмены контекста
	ctx5, cancel5 := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx5.Done():
				fmt.Println(5, "stopped")
				return
			default:
				fmt.Println(5, "working")
			}
			time.Sleep(time.Second)
		}
	}()

	cancel5()
	wg.Wait()

	var stop6 int32

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			if atomic.LoadInt32(&stop6) == 42 {
				fmt.Println(6, "stopped")
				return
			}
			fmt.Println(6, "working")
			time.Sleep(time.Second)
		}
	}()

	// Завершим рутину с помощью флага
	atomic.StoreInt32(&stop6, 42)

	wg.Wait()

	go func() {
		for {
			fmt.Println(8, "working")
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second)

	// Завершим рутину выходом из програмы
}
