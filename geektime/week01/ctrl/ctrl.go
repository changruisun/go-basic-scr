package main

import "fmt"

type User struct {
	name string
}

func loopMap() {
	m := make(map[string]*User)
	users := []User{
		{name: "tom"},
		{"jerry"},
	}
	// u 在栈里面，只分配了一次，循环遍历的时候，u的值每次都更改，因为只分配了一次地址没变
	for _, u := range users {
		fmt.Println(&u)
		m[u.name] = &u
	}

	for name, uP := range m {
		fmt.Println(name, uP.name)
	}

}
func main() {

	loopMap()

}
