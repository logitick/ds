package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := New()
	stack.Push(1)
	assert.Equal(t, 1, stack.top.value)
	assert.Nil(t, stack.top.next)

	stack.Push(5)
	assert.Equal(t, 5, stack.top.value)
	assert.Equal(t, 1, stack.top.next.value)

	stack.Push(2)
	assert.Equal(t, 2, stack.top.value)
	assert.Equal(t, 5, stack.top.next.value)
}

func TestStack_Len(t *testing.T) {
	stack := New()
	assert.Equal(t, 0, stack.Len())
	stack.len = 20
	assert.Equal(t, 20, stack.Len())
}

func TestStack_Push_incrementLen(t *testing.T) {
	stack := New()
	stack.Push(5)
	assert.Equal(t, 1, stack.Len())

	stack.Push(4)
	assert.Equal(t, 2, stack.Len())

	stack.Push(22)
	assert.Equal(t, 3, stack.Len())
}

func TestStack_Peek(t *testing.T) {
	stack := New()

	assert.Nil(t, stack.Peek())

	stack.Push(1)
	assert.Equal(t, 1, stack.Peek())

	stack.Push(9)
	assert.Equal(t, 9, stack.Peek())

	stack.Push(35)
	assert.Equal(t, 35, stack.Peek())
}

func TestStack_Pop(t *testing.T) {
	elements := []string{"hello", "world", "i am", "going"}
	stack := New()
	for _, i := range elements {
		stack.Push(i)
	}

	assert.Equal(t, "going", stack.Pop())
	assert.Equal(t, "i am", stack.Pop())
	assert.Equal(t, "world", stack.Pop())
	assert.Equal(t, "hello", stack.Pop())
}

func TestStack_Pop_decrementLen(t *testing.T) {
	elements := []string{"omg", "i", "am", "coding", "go"}
	stack := New()
	for _, i := range elements {
		stack.Push(i)
	}

	stack.Pop()
	assert.Equal(t, 4, stack.Len())
	stack.Pop()
	assert.Equal(t, 3, stack.Len())
	stack.Pop()
	assert.Equal(t, 2, stack.Len())
	stack.Pop()
	assert.Equal(t, 1, stack.Len())
	stack.Pop()
	assert.Equal(t, 0, stack.Len())
}

func TestStack_Push_afterEmpty(t *testing.T) {
	stack := New()
	stack.Push("wat")
	stack.Push("dis")
	stack.Pop()
	stack.Pop()

	stack.Push("a")
	assert.Equal(t, 1, stack.Len())
	assert.Equal(t, "a", stack.Pop())
}
