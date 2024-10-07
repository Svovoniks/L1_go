package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type AsyncMap struct {
	mutex    sync.Mutex
	writeMap map[int]int
}

func (am *AsyncMap) write(key int, value int) {
	am.mutex.Lock()
	am.writeMap[key] = value
	am.mutex.Unlock()
}

func writer(ctx context.Context, mp *AsyncMap) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			mp.write(rand.Int(), rand.Int())
			fmt.Println("wrote to map")
		}
	}
}

func main() {
	mp := AsyncMap{
		writeMap: make(map[int]int),
	}
	ctx, cancel := context.WithCancel(context.Background())

	go writer(ctx, &mp)
	go writer(ctx, &mp)
	go writer(ctx, &mp)

	time.Sleep(time.Second)

	cancel()

}
