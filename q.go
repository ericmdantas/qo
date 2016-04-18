package q

type Q struct {
	list   []interface{}
	length int
}

func NewQ() *Q {
	return &Q{list: []interface{}{}, length: 0}
}

func (q *Q) Unshift(info interface{}) {
	q.list = append([]interface{}{info}, q.list...)
}

func (q *Q) AddAmount(am int) {
	for i := 0; i < am; i++ {
		q.list = append(q.list, i)
	}

	q.length = am
}

func (q *Q) Push(info interface{}) {
	q.list = append(q.list, info)
	q.length += 1
}

func (q *Q) Pop() {
	if q.length > 0 {
		q.list = q.list[:len(q.list)-1]
		q.length -= 1
	}
}

func (q *Q) Shift() {
	if q.length > 0 {
		q.list = q.list[1:]
		q.length -= 1
	}
}

func (q *Q) Remove(el interface{}) {
	for i, v := range q.list {
		if v == el {
			q.list = append(q.list[:i], q.list[i+1:]...)
		}
	}

	q.length = len(q.list)
}

func (q *Q) Clean() {
	q.list = []interface{}{}
	q.length = 0
}
