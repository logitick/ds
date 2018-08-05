package stack

type Stacker interface {
	Pop() interface{}
	Peek() interface{}
	Push(item interface{})
	Len() int
}

type Stack struct {
	top *node
	len int
}

type node struct {
	value interface{}
	next  *node
}

func New() *Stack {
	return &Stack{top: nil, len: 0}
}

func (s *Stack) Peek() interface{} {
	if s.top == nil {
		return nil
	}
	return s.top.value
}

func (s *Stack) Push(item interface{}) {
	s.top = &node{value: item, next: s.top}
	s.len++
}

func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	n := s.top
	s.top = n.next
	s.len--
	return n.value
}

func (s *Stack) Len() int {
	return s.len
}
