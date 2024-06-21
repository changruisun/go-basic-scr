package main

type List[T any] interface {
	Add(index int, val T)
	Append(val T)
	Delete(index int)
}

type LinkList[T any] struct {
	head *node[T]
}

func (r *LinkList[T]) Add(index int, val T) {

}

type node[T any] struct {
	data T
}

func useList() {
	l := &LinkList[int]{}
	l.Add(1, 123)
	//l.Add(2, "123")

	l2 := &LinkList[string]{}
	//l2.Add(1, 123)
	l2.Add(2, "123")
}
