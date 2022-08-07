package linkedlist

import (
	"testing"
)

type linkedTestNodes struct {
	in   []int
	out  []int
	size int
}

func TestLinkedList_Delete(t *testing.T) {
	var node = linkedTestNodes{
		in:   []int{0, 1, 5, 3},
		out:  []int{0, 1, 5},
		size: 3,
	}

	var list = NewLinkedList()
	for _, i := range node.in {
		list.AddBySort(i)
	}

	list.Delete(3)
	if list.size != 3 {
		t.Errorf("size error, expect %d, but got %d", 3, list.size)
	}

	header := list.header
	for j := 0; j < len(node.out); j++ {
		if header.next != nil {
			if header.next.data != node.out[j] {
				t.Error("expect {} but got {}", node.out[j], header.next.data)
			}
			header = header.next
		}
	}

	list.Print()
}

func TestLinkedList_Search(t *testing.T) {
	var node = linkedTestNodes{
		in:   []int{0, 1, 5, 3},
		out:  []int{0, 1, 3, 5},
		size: 4,
	}

	var list = NewLinkedList()
	for _, i := range node.in {
		list.AddBySort(i)
	}

	data := 3
	if list.Search(data) != true {
		t.Errorf("expect %v, but got %v", true, false)
	}

}

func TestLinkedList_Print(t *testing.T) {
	var node = linkedTestNodes{
		in:   []int{0, 1, 5, 3},
		out:  []int{0, 1, 3, 5},
		size: 4,
	}

	var list = NewLinkedList()
	for _, i := range node.in {
		list.AddBySort(i)
	}

	list.Print()
}

func TestLinkedList_AddBySort(t *testing.T) {
	var nodes1 = linkedTestNodes{
		in:   []int{0, 1, 5, 3},
		out:  []int{0, 1, 3, 5},
		size: 4,
	}

	var nodes2 = linkedTestNodes{
		in:   []int{0, 0, 0, 0},
		out:  []int{0},
		size: 1,
	}

	var nodes3 = linkedTestNodes{
		in:   []int{0, 1, 1, 0},
		out:  []int{0, 1},
		size: 2,
	}

	nodeArr := []linkedTestNodes{nodes1, nodes2, nodes3}

	for _, node := range nodeArr {
		var list = NewLinkedList()
		for _, i := range node.in {
			list.AddBySort(i)
		}

		header := list.header
		if list.size != node.size {
			t.Errorf("size error, expect %d, but got %d", node.size, list.size)
		}

		for j := 0; j < len(node.out); j++ {
			if header.next != nil {
				if header.next.data != node.out[j] {
					t.Error("expect {} but got {}", node.out[j], header.next.data)
				}
				header = header.next
			}
		}

	}
}
