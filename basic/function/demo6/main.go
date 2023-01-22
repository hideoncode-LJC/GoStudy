package main

import "fmt"

type calculation func(int, int) int

func cal(x, y int, c calculation) int {
	return c(x, y)
}

func function(str string) calculation {
	if str == "+" {
		return add
	} else if str == "-" {
		return sub
	} else {
		return nil
	}
}

func add(x, y int) (z int) {
	z = x + y
	return z
}

func sub(x, y int) (z int) {
	z = x - y
	return z
}

func main() {
	ret1 := cal(1, 2, add)
	ret2 := cal(1, 2, sub)
	ret3 := function("+")(10, 20)
	ret4 := function("-")(10, 20)
	fmt.Println(ret1, ret2, ret3, ret4)
}
