package main

import "fmt"

const (
	n1 = iota
	n2
	n3
	n4 = 100
	n5
)

func main() {
	fmt.Println(n1, n2, n3, n4, n5)
}
