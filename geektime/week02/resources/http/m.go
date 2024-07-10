package main

import (
	"fmt"
	"time"
)

func main() {

	a := "2012-01-01"
	v, _ := time.Parse(time.DateOnly, a)
	fmt.Println(v.UnixMilli())
	print(v.Format(time.DateOnly))

}
