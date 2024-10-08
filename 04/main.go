package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func Worker(idx int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		val, ok := <-dataChan

		if !ok {
			return
		}

		fmt.Println(idx, val)
	}
}

// Я решил завершать работу воркеров закрывая канал с данными и затем
// ожидая их завершения с помощью WaitGroup.
// Я выбрал такой подход потому что после простого закрытия канала без ожидания
// некоторые воркеры могли бы до сих пор обрабатывать прочитанные данные что
// могло бы привести к тому что не все данные были бы веведенны. А использование
// Context мне показалось избыточным в данном случае.
func main() {
	var workerCount int
	_, err := fmt.Scan(&workerCount)
	if err != nil {
		fmt.Println("Couldn't parse workerCount")
		return
	}

	if workerCount < 1 {
		fmt.Println("Need at least 1 worker")
		return
	}

	dataChan := make(chan int)
	var wg sync.WaitGroup

	for i := range workerCount {
		wg.Add(1)
		go Worker(i, dataChan, &wg)
	}

	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt)

	counter := 0

	for {
		if len(exitChan) != 0 {
			break
		}

		dataChan <- counter
		counter += 1
	}

	close(dataChan)
	wg.Wait()
}
