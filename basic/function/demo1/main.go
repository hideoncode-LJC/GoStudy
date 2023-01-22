package main

import "fmt"

func getSum(x, y int) (int, int) {
	return x + y, 1
}

func sayHello() {
	fmt.Println("hello")
}

func main() {
	sayHello()
	fmt.Println("请输入两个数字")
	var x, y int
	fmt.Scanf("%d%d", &x, &y)
	res, _ := getSum(x, y)
	fmt.Println("两数之和为", res)
}
