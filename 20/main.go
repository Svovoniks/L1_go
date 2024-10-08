package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func reverseString(str string) string {
	sl := strings.Split(str, " ")
	slices.Reverse(sl)
	return strings.Join(sl, " ")
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	if sc.Err() != nil {
		fmt.Println("Couldn't read input")
	}
	st := sc.Text()

	fmt.Println(reverseString(st))
}
