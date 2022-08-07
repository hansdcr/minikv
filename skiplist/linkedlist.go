package skiplist

import (
	"bytes"
	"fmt"
)

const (
	defaultMaxHeight = 48
)

type Entry struct {
	Key []byte
	Val []byte
}

type Node struct {
	entry *Entry
	next  *Node
}

func NewNode(data *Entry) *Node {
	return &Node{entry: data, next: nil}
}

type LinkedList struct {
	header *Node
	size   int
}

func NewLinkedList() *LinkedList {
	header := &Node{
		entry: &Entry{},
		next:  nil,
	}

	return &LinkedList{
		header: header,
		size:   0,
	}
}

func (linkedlist *LinkedList) Add(entry *Entry) {
	newNode := NewNode(entry)
	prev := linkedlist.header

	for prev.next != nil {
		if Compare(entry.Key, prev.next.entry) > 0 {
			prev = prev.next
		} else if Compare(entry.Key, prev.next.entry) == 0 {
			prev.next.entry.Val = entry.Val //key相等,更新值
			return
		} else if Compare(entry.Key, prev.next.entry) < 0 {
			tmp := prev.next
			prev.next = newNode
			newNode.next = tmp
			linkedlist.size++
			return
		}
	}

	// 尾节点为空的情况
	prev.next = newNode
	linkedlist.size++
}

func (linkedlist *LinkedList) Search(key []byte) bool {
	header := linkedlist.header
	for header.next != nil {
		if Compare(key, header.entry) == 0 {
			return true
		}
		header = header.next
	}
	return false
}

func (linkedlist *LinkedList) Print() {
	header := linkedlist.header
	str := "nil"
	for header.next != nil {
		str = str + "--->" + string(header.next.entry.Key)
		header = header.next
	}
	fmt.Printf(str)
}

func (linkedlist *LinkedList) Delete(key []byte) {
	prev := linkedlist.header

	// 0 --> 1---> 2 --->5
	for prev.next != nil {
		if Compare(key, prev.next.entry) == 0 {
			tmp := prev.next.next
			prev.next = tmp
			linkedlist.size--
			return
		}
		prev = prev.next
	}
}

func Compare(key []byte, entry *Entry) int {
	return bytes.Compare(key, entry.Key)
}
