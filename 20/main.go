package main

import (
	"bytes"
	"fmt"
	"strings"
)

func reverseString(str string) string {
	buffer := new(bytes.Buffer)

	sl := strings.Split(str, " ")
	for i := range sl {
		buffer.WriteString(sl[len(sl) - i - 1])
        buffer.WriteRune(' ')
	}

	return buffer.String()
}

func main() {
	st := "snow dog sun"

	fmt.Println(reverseString(st))
}
