package main

import "fmt"

func binSearch(arr []int, target int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		if guess == target {
			return mid
		} else if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	return -1
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binSearch(arr, 5))
	fmt.Println(binSearch(arr, 15))
	fmt.Println(binSearch(arr, 1))
}
