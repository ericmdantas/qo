package main

import (
	"fmt"
	"time"
)

const (
	max int = 10000000
)

type qo struct {
	list []int
}

func (q *qo) push(i int) {
	q.list = append(q.list, i)
}

func (q *qo) pop() {
	q.list = q.list[:len(q.list)-1]
}

func main() {
	q := qo{}

	queue := q.list

	start := time.Now()
	for i := 0; i < max; i++ {
		q.push(i)
	}

	elapsed := time.Since(start)

	fmt.Println(elapsed)

	start = time.Now()

	for i := 0; i < max; i++ {
		q.pop()
	}

	elapsed = time.Since(start)

	fmt.Println(elapsed)
	fmt.Println(queue)
}
