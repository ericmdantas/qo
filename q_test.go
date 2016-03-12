package q

import (
	"testing"
)

/*func BenchmarkUnshift(b *testing.B) {
	q := NewQ()

	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}*/

func BenchmarkPush(b *testing.B) {
	q := NewQ()

	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	q := NewQ()

	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkShift(b *testing.B) {
	q := NewQ()

	for i := 0; i < b.N; i++ {
		q.Shift()
	}
}

func TestPush(t *testing.T) {
	q := NewQ()

	if len(q.list) != 0 {
		t.Error("expected 0, got:", len(q.list))
	}

	q.Push(1)

	if len(q.list) != 1 {
		t.Error("expected 1, got:", len(q.list))
	}

	q.Push(2)

	if len(q.list) != 2 {
		t.Error("expected 2, got:", len(q.list))
	}
}
