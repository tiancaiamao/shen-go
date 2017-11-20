package vm

import (
	"fmt"

	"github.com/tiancaiamao/shen-go/kl"
)

type queue struct {
	data []kl.Obj
	head int
	tail int
}

func (q *queue) init(count int) {
	q.data = make([]kl.Obj, count)
	q.head = 0
	q.tail = 0
}

func (q *queue) count() int {
	return (q.tail + len(q.data) - q.head) % len(q.data)
}

func (q *queue) empty() bool {
	return q.tail == q.head
}

func (q *queue) push(o kl.Obj) {
	nextPos := (q.tail + 1) % len(q.data)
	if nextPos == q.head {
		q.expand()
		nextPos = (q.tail + 1) % len(q.data)
	}
	q.data[q.tail] = o
	q.tail = nextPos
}

func (q *queue) pop() kl.Obj {
	if q.empty() {
		panic("pop a empty queue")
	}
	ret := q.data[q.head]
	q.head = (q.head + 1) % len(q.data)
	return ret
}

func (q *queue) expand() {
	tmp := make([]kl.Obj, len(q.data)*2)
	i := 0
	for pos := q.head; pos != q.tail; pos = (pos + 1) % len(q.data) {
		tmp[i] = q.data[pos]
		i++
	}

	q.data = tmp
	q.head = 0
	q.tail = i
}

func (q *queue) debug() {
	fmt.Println("...", q.head, q.tail, len(q.data))
}
