package main

import "fmt"

type myStruct struct {
	a int8
	b int16
	c float64
}

func main() {
	var a myStruct
	a.a = 1
	a.b = 2
	a.c = 1.1111111
	fmt.Println(&a.a)
	fmt.Println(&a.b)
	fmt.Println(&a.c)

}
