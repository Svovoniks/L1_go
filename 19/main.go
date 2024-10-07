package main

import (
	"fmt"
	"slices"
)

func reverse(str string) string {
	runes := []rune(str)

	slices.Reverse(runes)
	return string(runes)
}

func main() {
	st := "главрыба 🌞"

	fmt.Println(reverse(st))
}
