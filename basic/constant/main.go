package main

import "fmt"

// 普通声明
const PI = 3.1415
const e = 2.7

//批量声明
const (
	a = 1
	b = 2
	c = 3
)

//iota 的使用
const (
	n1 = iota
	n2
	_
	n3
)
const n5 = iota

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	x1, x2 = iota + 1, iota + 2 //1,2
	x3, x4                      //2,3
	x5, x6                      //3,4
)

func main() {
	fmt.Println(PI, e)
	fmt.Println(n1, n2, n3, n5)
	fmt.Println(KB, MB)
	fmt.Println(x1, x2, x3, x4, x5, x6)
}
