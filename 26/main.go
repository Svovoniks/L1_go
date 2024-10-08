package main

import (
	"fmt"
	"strings"
)

func IsUnique(str string) bool {
	str = strings.ToLower(str)
	mp := make(map[rune]bool)

	for _, i := range []rune(str) {
		if mp[i] {
			return false
		}
		mp[i] = true
	}

	return true
}

func main() {
	var str string
	if _, err := fmt.Scanln(&str); err != nil {
		fmt.Println("Couldn't parse input")
		return
	}
	fmt.Println(IsUnique(str))

}
