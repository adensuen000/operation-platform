package main

import "fmt"

func main() {
	type tt struct {
		Name string
	}
	var TT *tt
	TT = new(tt)
	TT.Name = "test"
	fmt.Println("打印&TT: ", TT)
}
