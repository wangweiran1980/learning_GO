package fib

import "testing"

func TestFib(t *testing.T) {
	f := Fibonacci()
	t.Log(f())
	t.Log(f())
	t.Log(f())
	t.Log(f())
	t.Log(f())
	t.Log(f())
	printFileContents(f)
}
