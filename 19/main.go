package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func reverse(str string) string {
	runes := []rune(str)

	slices.Reverse(runes)
	return string(runes)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	if sc.Err() != nil {
		fmt.Println("Couldn't read input")
        return
	}
	st := sc.Text()

	fmt.Println(reverse(st))
}
