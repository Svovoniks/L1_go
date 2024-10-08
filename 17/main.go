package main

import "fmt"

func BinSearch(val int, arr []int) int {
	var end int = len(arr) - 1
	var start int = 0

	for start <= end {
		var mid int = start + (end-start)/2

		if arr[mid] == val {
			return mid
		}

		if arr[mid] < val {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}

	return -1

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, i := range arr {
		fmt.Println(i, BinSearch(i, arr))
	}

	fmt.Println(0, BinSearch(10, arr))
	fmt.Println(10, BinSearch(10, arr))
}
