package main

import "fmt"

func intersect(arr1 *[]int, arr2 *[]int) []int {
	mp := make(map[int]bool)
	var res []int

	for _, i := range *arr1 {
		mp[i] = true
	}

	for _, i := range *arr2 {
		if mp[i] {
			res = append(res, i)
		}
	}

	return res

}

func main() {
	set1 := []int{20, 2, 3, 4, 5}
	set2 := []int{2, 3, 5, 7, 8, 9, 20}

	fmt.Println(intersect(&set1, &set2))
}
