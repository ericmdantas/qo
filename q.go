package q

const (
	max int = 10000000
)

type Q struct {
	list []int
}

/*func (q *Q) Unshift(i int) {
	q.list = append(q, q...)
}*/

func NewQ() *Q {
	return &Q{[]int{}}
}

func (q *Q) Push(i int) {
	q.list = append(q.list, i)
}

func (q *Q) Pop() {
	if len(q.list) > 0 {
		q.list = q.list[:len(q.list)-1]
	}
}

func (q *Q) Shift() {
	if len(q.list) > 0 {
		q.list = q.list[1:]
	}
}

func (q *Q) Remove(el int) {
	for i, v := range q.list {
		if v == el {
			q.list = append(q.list[:i], q.list[i+1:]...)
		}
	}
}
