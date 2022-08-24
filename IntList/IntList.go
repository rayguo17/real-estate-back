package IntList

import "fmt"

type IntList struct {
	item int
	next *IntList
}

func NewIntList(x int) *IntList {
	return &IntList{
		item: x,
	}
}

func (l *IntList) AddLast(x int) {
	tmp := l
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &IntList{item: x}
}

func (l *IntList) Print() {
	tmp := l
	for tmp != nil {
		fmt.Printf("%d ", tmp.item)
		tmp = tmp.next
	}
	fmt.Println("")
}
