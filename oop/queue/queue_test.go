package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.IsEmpty())
	t.Log(q.Pop())
	t.Log(q.IsEmpty())
}
