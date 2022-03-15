package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	for i := range arr2 {
		fmt.Println(arr2[i])
	}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}
	printArray(arr1)
}
