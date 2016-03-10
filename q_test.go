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
