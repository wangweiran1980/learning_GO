package adder

import (
	"testing"
)

func TestAdder(t *testing.T) {
	// a := Adder()
	a := Adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		t.Log(s)
	}
}
