package main

import "fmt"

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	var c calculation = add
	ret1 := c(1, 2)
	c = sub
	ret2 := c(1, 2)
	fmt.Println(ret1, ret2)
}
