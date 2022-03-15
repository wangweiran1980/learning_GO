package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func main() {
	// var s1 []int //nil

	// for i := 0; i < 100; i++ {
	// 	printSlice(s1)
	// 	s1 = append(s1, 2*i+1)
	// }

	s2 := make([]int, 16, 32)
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println("删除元素...")
	s2[0] = 1
	s2[1] = 2
	s2[3] = 4
	s2[4] = 6
	fmt.Println(s2)
	s2 = append(s2[:2], s2[3:5]...)
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println("Popping from front...")
	s2 = s2[1:]
	fmt.Println(s2)
	fmt.Println("poping from back...")
	s2 = s2[:len(s2)-1]
	fmt.Println(s2)
}
