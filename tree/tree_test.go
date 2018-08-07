package tree

import "testing"

func TestBinarySearchTree_Size(t *testing.T) {
	type fields struct {
		root *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &BinarySearchTree{
				root: tt.fields.root,
			}
			if got := t.Size(); got != tt.want {
				t.Errorf("BinarySearchTree.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
