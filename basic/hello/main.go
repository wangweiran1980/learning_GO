package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/cmplx"
	"os"
	"strconv"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d,%q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 5, 6, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 7, 8, true, "def"
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func emums() {
	const (
		cpp = iota
		java
		python
		javascript
		// php
		_
		golang
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
	fmt.Println(cpp, java, python, javascript, golang)
}

func readFile(filename string) {
	if context, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(context))
	}
}

func readFile2(filename string) {
	f, err := os.Open(filename)
	if err == io.EOF {
		return
	}
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func grade(score int) (string, error) {
	g := ""
	switch {
	case score < 0 || score > 100:
		return g, fmt.Errorf("错误的成绩: %d", score)
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g, nil
}

func convertToBin(n int) (res string) {
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 函数可以作为参数传递
func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println("Hello, GO")
	// variableZeroValue()
	// variableInitialValue()
	// variableTypeDeduction()
	// variableShorter()
	// fmt.Println(aa, ss, bb)
	// euler()
	// triangle()
	// consts()
	// emums()
	// readFile("基础语法/hello/main.go")
	// fmt.Println(grade(90))
	// fmt.Println(convertToBin(5), convertToBin(13))
	// readFile2("基础语法/hello/main.go")
	// a, _ := div(3, 4)
	// fmt.Println(a)
	// fmt.Println(apply(func(a, b int) int {
	// 	return int(math.Pow(float64(a), float64(b)))
	// }, 3, 4))
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
