package main

import "fmt"

type List[T any] interface {
	Delete(idx int)
}

type ArrayList[T any] struct {
	Vals []Node[T]
}

type Node[T any] struct {
	Data T
}

func (r *ArrayList[T]) Delete(idx int) {
	if idx < 0 || idx > len(r.Vals) {
		panic("idx args error")
	}
	// 这个地方不知道如何控制len的属性，所以重新生成了切片
	newVals := make([]Node[T], len(r.Vals)-1, len(r.Vals))
	for i := 0; i < idx; i++ {
		newVals[i] = r.Vals[i]
	}
	for i := idx; i < len(r.Vals)-1; i++ {
		newVals[i] = r.Vals[i+1]
	}
	//r.Vals[len(r.Vals)-1] = Node[T]{}
	r.Vals = newVals

	// todo 缩容, 重新生成切片

}

func main() {
	var l *ArrayList[int]
	l = &ArrayList[int]{Vals: []Node[int]{Node[int]{1}, Node[int]{2}, Node[int]{3}}}
	l.Delete(1)
	fmt.Printf("%v \n", l)
	l = &ArrayList[int]{Vals: []Node[int]{Node[int]{1}, Node[int]{2}, Node[int]{3}}}
	l.Delete(2)
	fmt.Printf("%v \n", l)
	l = &ArrayList[int]{Vals: []Node[int]{Node[int]{1}, Node[int]{2}, Node[int]{3}}}
	l.Delete(0)
	fmt.Printf("%v \n", l)
}
