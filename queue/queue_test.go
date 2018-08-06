package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Len(t *testing.T) {
	q := New()
	assert.Equal(t, 0, q.Len())

	q.len = 2
	assert.Equal(t, 2, q.Len())

	q.len = 23
	assert.Equal(t, 23, q.Len())
}
func TestQueue_Enq(t *testing.T) {
	q := New()

	q.Enq('a')
	assert.Equal(t, 'a', q.first.value)
	assert.Nil(t, q.first.next)

	assert.Equal(t, 'a', q.last.value)
	assert.Nil(t, q.last.next)

	q.Enq('b')
	assert.Equal(t, 'a', q.first.value)
	assert.Equal(t, 'b', q.first.next.value)
	assert.Equal(t, 'b', q.last.value)
}

func TestQueue_Enq_incLen(t *testing.T) {
	q := New()
	assert.Equal(t, 0, q.Len())

	q.Enq('a')
	assert.Equal(t, 1, q.Len())

	q.Enq('b')
	assert.Equal(t, 2, q.Len())

	q.Enq('c')
	assert.Equal(t, 3, q.Len())
}

func TestQueue_Deq(t *testing.T) {
	q := New()
	q.Enq("hello")
	q.Enq("world")
	q.Enq("あ")

	assert.Equal(t, "hello", q.Deq())
	assert.Equal(t, "world", q.Deq())
	assert.Equal(t, "あ", q.Deq())
}

func TestQueue_Deq_DecLen(t *testing.T) {
	q := New()
	q.Enq('a')
	q.Enq('b')
	q.Enq('c')

	q.Deq()
	assert.Equal(t, 2, q.Len())

	q.Deq()
	assert.Equal(t, 1, q.Len())

	q.Deq()
	assert.Equal(t, 0, q.Len())
}
