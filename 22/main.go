package main

import (
	"fmt"
	"math/big"
)

func main() {
	var fl1Str string
	var fl2Str string

	if _, err := fmt.Scan(&fl1Str); err != nil {
		fmt.Println("Couldn't parse input")
		return
	}

	if _, err := fmt.Scan(&fl2Str); err != nil {
		fmt.Println("Couldn't parse input")
		return
	}
	fl1, _, err1 := big.ParseFloat(fl1Str, 10, 10, big.ToNearestAway)
	if err1 != nil {
		fmt.Println("Couldn't parse input")
		return
	}

	fl2, _, err2 := big.ParseFloat(fl2Str, 10, 10, big.ToNearestAway)

	if err2 != nil {
		fmt.Println("Couldn't parse input")
		return
	}

	res := &big.Float{}

	minSupported := big.NewFloat(1 << 20)

	if fl1.Cmp(minSupported) != 1 || fl2.Cmp(minSupported) != 1 {
		fmt.Println("Number too small")
		return
	}

	res.Add(fl1, fl2)
	fmt.Println("Added", res.Text('f', 3))

	res.Sub(fl1, fl2)
	fmt.Println("Subtracted", res.Text('f', 3))

	res.Mul(fl1, fl2)
	fmt.Println("Multiplied", res.Text('f', 3))

	res.Quo(fl1, fl2)
	fmt.Println("Divided", res.Text('f', 33))
}
