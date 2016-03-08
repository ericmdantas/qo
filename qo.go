package main

import (
	"fmt"
	"time"
)

const (
	MAX int = 9999999
)

type qo struct {
	list []interface{}
}

func (q *qo) push(d interface{}) {
	q.list = append(q.list, d)
}

func (q *qo) unshift(d interface{}) {
	q.list = append(q.list, d)
}

func (q *qo) pop() {
	q.list = q.list[1:]
}

func (q *qo) shift() {
	q.list = q.list[:len(q.list)-1]
}

func logAdd(s string, cb func(info interface{}), list []interface{}) {
	t := time.Now()

	fmt.Println(s)

	for i := 0; i < MAX; i++ {
		cb(i)
	}

	fmt.Println(time.Since(t))
	fmt.Println(len(list))
}

func logRemove(s string, cb func(), list []interface{}) {
	t := time.Now()

	fmt.Println(s)

	for i := 0; i < MAX; i++ {
		cb()
	}

	fmt.Println(time.Since(t))
	fmt.Println(len(list))
}

func main() {
	q := qo{}

	logAdd("push", q.push, q.list)
	logAdd("unshift", q.unshift, q.list)

	logRemove("pop", q.pop, q.list)
	logRemove("shift", q.pop, q.list)
}
