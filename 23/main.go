package main

import "fmt"

func RemoveElementAt(arr []int, idx int) []int {
	newArr := make([]int, 0)
    newArr = append(newArr, arr[:idx]...)
    newArr = append(newArr, arr[idx+1:]...)
	return newArr
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	for i := range len(arr) {
		narr := RemoveElementAt(arr, i)
		fmt.Println(narr)
	}

}
