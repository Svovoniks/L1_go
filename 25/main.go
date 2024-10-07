package main

import (
	"time"
)

func sleep(sec int) {
	<-time.After(time.Second * time.Duration(sec))
}

func main() {
	sleep(10)
}
