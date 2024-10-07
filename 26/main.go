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

func main(){
    var str string
    fmt.Scanln(&str)
    fmt.Println(IsUnique(str))

}
