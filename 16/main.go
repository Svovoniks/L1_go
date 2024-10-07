package main

import "fmt"

func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}

	var pivIdx int = len(arr) / 2
	piv := arr[pivIdx]
	arr[pivIdx], arr[len(arr)-1] = arr[len(arr)-1], arr[pivIdx]

	lessIdx := 0

	for i := range arr {
		if arr[i] < piv {
			arr[lessIdx], arr[i] = arr[i], arr[lessIdx]
			lessIdx++
		}
	}

	arr[lessIdx], arr[len(arr)-1] = arr[len(arr)-1], arr[lessIdx]

	quicksort(arr[:lessIdx])
	quicksort(arr[lessIdx+1:])
}

func main() {
	arr := []int{13, 2, 93, 34, 29295, 1}
	quicksort(arr)
	fmt.Println(arr)
}
