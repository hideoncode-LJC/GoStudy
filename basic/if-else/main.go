package main

import (
	"fmt"
	_ "net"
)

func ifDemoOne(x int) string {
	res := ""
	if x%2 == 0 {
		res = res + "偶数"
	} else if x%2 == 1 {
		res = res + "奇数"
	} else {
		res = res + "whatever"
	}
	return res
}

func ifDemoTwo(x int) string {
	res := ""
	if score := x; score%2 == 0 {
		res = res + "even"
	} else if x%2 == 1 {
		res = res + "odd"
	} else {
		res = res + "whatever"
	}
	return res
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(ifDemoOne(number))
	fmt.Println(ifDemoTwo(number))
}
