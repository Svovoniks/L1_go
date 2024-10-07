package main

import (
	"fmt"
)

func main() {
	var num int64
	var i int8
	var bitVal int8

	if _, err := fmt.Scan(&num, &i, &bitVal); err != nil ||
		!(i >= 0 && i < 64) || (bitVal != 0 && bitVal != 1) {
		fmt.Println("i need int64, i(0 <= i < 64), bit value(1 or 0)")
		return
	}

    fmt.Printf("num before bin: %b dec: %v\n", num, num)
	fmt.Printf("i %b\n", i)

	if bitVal == 1 {
		var mask int64 = (1 << i)
		num |= mask
	} else {
		var mask int64 = ^(1 << i)
		num &= mask
	}

    fmt.Printf("num after bin: %b, dec: %v\n", num, num)
}
