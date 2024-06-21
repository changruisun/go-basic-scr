package main

import "fmt"

func Sum[T Number](vals []T) T {
	var res T
	for _, v := range vals {
		res += v
	}
	return res

}

func Max[T Number](vals []T) T {
	t := vals[0]
	for i := 1; i < len(vals); i++ {
		if t < vals[i] {
			t = vals[i]
		}
	}
	return t
}

func Min[T Number](vals []T) T {
	t := vals[0]
	for i := 1; i < len(vals); i++ {
		if t > vals[i] {
			t = vals[i]
		}
	}
	return t
}

func Find[T any](vals []T, filter func(val T) bool) T {
	for _, v := range vals {
		if filter(v) {
			return v
		}
	}
	var t T
	return t
}

func Insert[T Number](idx int, val T, vals []T) []T {
	if idx < 0 || idx > len(vals) {
		panic("index error")
	}

	vals = append(vals, val)
	for i := len(vals) - 1; i > idx; i-- {
		if i-1 >= 0 {
			vals[i] = vals[i-1]
		}
	}
	vals[idx] = val
	return vals
}

type Number interface {
	~int | uint | int32
}

type Integer int

func UseSum() {
	fmt.Println(Sum([]Integer{1, 2, 3}))

	fmt.Printf("%v \n", Insert(3, 12, []Integer{1, 2, 3}))
	fmt.Printf("%v \n", Insert(2, 12, []Integer{1, 2, 3}))
	fmt.Printf("%v \n", Insert(0, 12, []Integer{1, 2, 3}))
	fmt.Printf("%v \n", Insert(1, 12, []Integer{1, 2, 3}))
	//fmt.Printf("%v \n", Insert(12, 12, []Integer{1, 2, 3}))
}
