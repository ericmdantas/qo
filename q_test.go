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

func TestPushLength(t *testing.T) {
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

func TestPushInfo(t *testing.T) {
	q := NewQ()

	q.Push(1)

	if q.list[0] != 1 {
		t.Error("expected first position to be 1, but it was: ", q.list[0])
	}

	q.Push(2)

	if q.list[1] != 2 {
		t.Error("expected second position to be 2, but it was: ", q.list[1])
	}

	q.Push(9999)

	if q.list[2] != 9999 {
		t.Error("expected third position to be 9999 , but it was: ", q.list[2])
	}
}

func TestPopLength(t *testing.T) {
	q := NewQ()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	q.Pop()

	if len(q.list) != 2 {
		t.Error("expected 2, got: ", len(q.list))
	}

	q.Pop()

	if len(q.list) != 1 {
		t.Error("expected 1, got: ", len(q.list))
	}

	q.Pop()

	if len(q.list) != 0 {
		t.Error("expected 0, got: ", len(q.list))
	}

	q.Pop()

	if len(q.list) != 0 {
		t.Error("expected 0, got: ", len(q.list))
	}
}

func TestPopInfo(t *testing.T) {
	q := NewQ()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	q.Pop()

	if q.list[1] != 2 {
		t.Error("expected last position to be 2, but got: ", q.list[1])
	}

	q.Pop()

	if q.list[0] != 1 {
		t.Error("expected last position be 1, but got: ", q.list[0])
	}
}

func TestShiftLength(t *testing.T) {
	q := NewQ()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	q.Shift()

	if len(q.list) != 2 {
		t.Error("expected 2, got: ", len(q.list))
	}

	q.Shift()

	if len(q.list) != 1 {
		t.Error("expected 1, got: ", len(q.list))
	}

	q.Shift()

	if len(q.list) != 0 {
		t.Error("expected 0, got: ", len(q.list))
	}

	q.Shift()

	if len(q.list) != 0 {
		t.Error("expected 0, got: ", len(q.list))
	}
}

func TestShiftInfo(t *testing.T) {
	q := NewQ()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	q.Shift()

	if q.list[0] != 2 {
		t.Error("expected first position to be 2, but got: ", q.list[0])
	}

	q.Shift()

	if q.list[0] != 3 {
		t.Error("expected first position be 3, but got: ", q.list[0])
	}
}
