package skiplist

import (
	"bytes"
	"fmt"
	"strconv"
)

const (
	defaultMaxLevel = 2
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

type SkipList struct {
	levels   []*Node
	maxLevel int
	size     int64
}

func NewSkipList() *SkipList {
	return &SkipList{
		levels:   make([]*Node, defaultMaxLevel+1),
		maxLevel: defaultMaxLevel,
		size:     0,
	}
}

func (skiplist *SkipList) Add(entry *Entry) {
	maxLevel := skiplist.maxLevel

	// 遍历每一层
	for i := maxLevel; i >= 0; i-- {
		// 遍历每一层的链表

		newNode := NewNode(entry)

		if skiplist.levels[i] == nil {
			skiplist.levels[i] = newNode
			continue
		}

		prev := skiplist.levels[i]
		for prev.next != nil {
			if skiplist.Compare(entry.Key, prev.entry) > 0 {
				tmp := prev.next
				prev.next = newNode
				newNode.next = tmp
				break
			} else if skiplist.Compare(entry.Key, prev.entry) == 0 {
				prev.entry.Val = entry.Val
				break
			} else if skiplist.Compare(entry.Key, prev.entry) < 0 {
				prev = prev.next
			}
		}
		prev.next = newNode
	}
}
func (skiplist *SkipList) Print() {
	maxLevel := skiplist.maxLevel
	// 遍历每一层
	for i := maxLevel; i >= 0; i-- {
		// 遍历每一层的链表
		header := skiplist.levels[i]
		str := strconv.Itoa(i) + " nil-->"
		for header != nil {
			str = str + string(header.entry.Key) + "-->"
			header = header.next
		}
		fmt.Printf(str + "\n")
	}
}

func (skiplist *SkipList) Compare(key []byte, entry *Entry) int {
	return bytes.Compare(key, entry.Key)
}
