package queue

type Queuer interface {
	Len() int
	Enq(interface{})
	Deq() interface{}
}

type (
	node struct {
		value interface{}
		next  *node
	}
	Queue struct {
		first *node
		last  *node
		len   int
	}
)

func New() *Queue {
	return &Queue{first: nil, len: 0}
}

func (q *Queue) Len() int {
	return q.len
}

// Enq enqueues a value to the end of the queue
func (q *Queue) Enq(v interface{}) {
	n := &node{value: v, next: nil}
	if q.last != nil {
		q.last.next = n
	}
	q.last = n
	q.len++

	if q.first == nil {
		q.first = n
	}

	if q.first != n && q.first.next == nil {
		q.first.next = n
	}
}

// Deq dequeues the first item from the queue
func (q *Queue) Deq() interface{} {
	if q.first == nil {
		return nil
	}
	n := q.first
	q.first = n.next
	q.len--

	return n.value
}
