package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mutex sync.RWMutex
}

func (c *Counter) AddOne() {
	c.mutex.Lock()
	c.count++
	c.mutex.Unlock()
}

func (c *Counter) GetCount() int {
	defer c.mutex.RUnlock()

	c.mutex.RLock()
	return c.count
}

func main() {
	counter := Counter{}
	wg := new(sync.WaitGroup)

	worker := func() {
		for range 10 {
			counter.AddOne()
		}
		wg.Done()
	}

	wg.Add(1)
	go worker()
	wg.Add(1)
	go worker()
	wg.Add(1)
	go worker()
	wg.Add(1)
	go worker()
	wg.Add(1)
	go worker()

	wg.Wait()
	fmt.Println(counter.GetCount())

}
