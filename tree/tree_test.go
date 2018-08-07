package tree

import (
	"testing"
)

func TestBinarySearchTree_Size(t *testing.T) {
	tests := []struct {
		name     string
		items    []interface{}
		expected int
	}{
		{"Size of 3", []interface{}{1, 2, 3}, 3},
		{"Size of 0", []interface{}{nil}, 0},
		{"Strings", []interface{}{"hellow", "world", "this", "is a", "test"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := New()
			for _, i := range tt.items {
				tree.Insert(i)
			}
			if got := tree.Size(); got != tt.expected {
				t.Errorf("BinarySearchTree.Size() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBinarySearchTree_Insert_root(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		args []interface{}
	}{
		{"one args", &Node{value: 5}, []interface{}{5}},
		{"multiple args", &Node{value: 1}, []interface{}{1, 2, 3}},
		{"no args", nil, []interface{}{nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := New()
			for _, i := range tt.args {
				tr.Insert(i)
			}
			if got := tr.root.value; got != tt.root {
				t.Errorf("BinarySearchTree.root() = %v, want %v", got, tt.root.value)
			}
		})
	}
}
