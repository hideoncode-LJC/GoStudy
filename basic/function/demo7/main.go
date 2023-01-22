package main

import "fmt"

type calculation func(int, int) int

func cal() (func(int, int) int, string) {

	add := func(x, y int) int {
		return x + y
	}
	fmt.Printf("%T\n", add)
	fmt.Println("两数之和为", add(10, 20))

	var c calculation = func(x, y int) int {
		return x * y
	}
	fmt.Printf("%T\n", c)
	fmt.Println("两数之积为", c(10, 20))

	answer := func(x, y int) int {
		return x / y
	}(20, 10)
	fmt.Println("两数之商", answer)

	return func(i int, i2 int) int {
		return i - i2
	}, "success"
}

func main() {
	var cc calculation
	cc, _ = cal()
	fmt.Println("两数之差为", cc(100, 10))
}
