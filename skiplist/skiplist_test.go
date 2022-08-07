package skiplist

import (
	"bytes"
	"testing"
)

func TestSkipList_Compare(t *testing.T) {
	slice1 := []byte{1}
	slice2 := []byte{0}

	if bytes.Compare(slice1, slice2) <= 0 {
		t.Errorf("Failed. expect 1,but got -1")
	}

	slice3 := []byte{'a', 'b', 'c'}
	slice4 := []byte{'A', 'B', 'C'}

	if bytes.Compare(slice3, slice4) <= 0 {
		t.Errorf("Failed. expect 1,but got -1")
	}

	slice5 := []byte{'a', 'b', 'c'}
	slice6 := []byte{'b', 'c', 'd'}

	if bytes.Compare(slice5, slice6) >= 0 {
		t.Errorf("Failed. expect -1,but got 1")
	}
}

func TestSkipList_Add(t *testing.T) {
	entry1 := &Entry{
		Key: []byte{'a', 'b', 'c'},
		Val: []byte{'a', 'b', 'c'},
	}

	entry2 := &Entry{
		Key: []byte{'a', 'b', 'e'},
		Val: []byte{'a', 'b', 'e'},
	}

	entry3 := &Entry{
		Key: []byte{'a', 'b', 'd'},
		Val: []byte{'a', 'b', 'd'},
	}

	skiplist := NewSkipList()
	skiplist.Add(entry1)
	skiplist.Add(entry2)
	skiplist.Add(entry3)

	skiplist.Print()
}
