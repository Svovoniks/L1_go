package main

import "fmt"

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	mp := make(map[int][]float32)

	for _, i := range arr {
		idx := int(i) / 10 * 10
		mp[idx] = append(mp[idx], i)
	}

	fmt.Println(mp)
}
