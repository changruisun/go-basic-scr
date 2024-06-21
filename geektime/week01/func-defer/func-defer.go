package main

import "fmt"

func DeferCloseV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Print(fmt.Sprintf("%d ", i))
		}()
	}
}

func DeferCloseV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Print(fmt.Sprintf("%d ", val))
		}(i)
	}
}

func DeferCloseV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Print(fmt.Sprintf("%d ", j))
		}()
	}
}

func main() {

	DeferCloseV1()
	fmt.Println()
	DeferCloseV2()
	fmt.Println()
	DeferCloseV3()

}
