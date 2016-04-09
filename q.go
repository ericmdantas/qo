package q

const (
	max int = 10000000
)

type Q struct {
	list   []int
	length int
}

func NewQ() *Q {
	return &Q{list: []int{}, length: 0}
}

func (q *Q) Unshift(i int) {
	q.list = append([]int{i}, q.list...)
}

func (q *Q) AddAmount(am int) {
	for i := 0; i < am; i++ {
		q.list = append(q.list, i)
	}

	q.length = len(q.list)
}

func (q *Q) Push(i int) {
	q.list = append(q.list, i)
	q.length = len(q.list)
}

func (q *Q) Pop() {
	if q.length > 0 {
		q.list = q.list[:len(q.list)-1]
	}

	q.length = len(q.list)
}

func (q *Q) Shift() {
	if q.length > 0 {
		q.list = q.list[1:]
	}

	q.length = len(q.list)
}

func (q *Q) Remove(el int) {
	for i, v := range q.list {
		if v == el {
			q.list = append(q.list[:i], q.list[i+1:]...)
		}
	}

	q.length = len(q.list)
}

func (q *Q) Clean() {
	q.list = []int{}
	q.length = 0
}
