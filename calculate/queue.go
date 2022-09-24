package calculate

type Node struct {
	Name     string
	Distance int
}

type Queue []Node

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(n Node) {
	*q = append(*q, n)
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Shift() Node {
	firstElement, updatedQueue := (*q)[0], (*q)[1:]
	*q = updatedQueue
	return firstElement
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
