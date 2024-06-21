package main

import "fmt"

type List interface {
	Add(index int, val any)
	Append(val any)
	Delete(index int)
}

type LinkList struct {
	head node
}

type node struct {
	data any
	next *node
}

func (l *LinkList) Add(index int, val any) {
	fmt.Printf("%d %v \n", index, val)

}

func userList() {
	l := &LinkList{}
	l.Add(1, 123)
	l.Add(2, "123")
	l.Add(3, nil)
}

func main() {
	userList()
}
