package cdawg

import "github.com/nthapaliya/wordgraph/dawg"

// Queue used internally by Compress()

func newQueue(size int) *queue {
	return &queue{
		nodes: make([]*dawg.State, size),
		size:  size,
	}
}

type queue struct {
	nodes []*dawg.State
	size  int
	head  int
	tail  int
	count int
}

func (q *queue) push(n *dawg.State) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*dawg.State, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func (q *queue) pop() *dawg.State {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}
