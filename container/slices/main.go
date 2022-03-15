package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])
	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println("s1 =", s1)
	fmt.Println("arr =", arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println("s2 =", s2)
	fmt.Println("arr =", arr)

	fmt.Println("Reslice...")
	fmt.Println("s2 =", s2)
	s2 = s2[:5]
	fmt.Println("s2 =", s2)
	s2 = s2[2:]
	fmt.Println("s2 =", s2)

	fmt.Println("额外...")
	fmt.Println("arr =", arr)
	s3 := arr[2:6]
	fmt.Println("s3 =", s3, len(s3), cap(s3))
	s4 := s3[3:5]
	fmt.Println("s4 =", s4, cap(s4))
	s4 = append(s4, 10)
	fmt.Println("s4 =", s4, cap(s4))
	s4 = append(s4, 11)
	fmt.Println("s4 =", s4, cap(s4))
	s4 = append(s4, []int{1, 2, 3}...)
	fmt.Println("s4 =", s4, cap(s4))
}
