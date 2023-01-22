package main

import "fmt"

func testA(a, b int) (x, y int) {
	return a, b
}

func main() {
	a, b := testA(1, 2)
	fmt.Println(a, b)
}
