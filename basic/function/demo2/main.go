package main

import (
	"fmt"
)

/*
可变参数
可变参数只能作为函数参数的最后一个
*/

func getSum(x int, y ...int) int {
	//x 是一个切片
	fmt.Println(y)
	sum := 0
	for _, v := range y {
		sum = sum + v
	}
	sum = sum + x
	return sum
}

func main() {
	ret1 := getSum(0)
	ret2 := getSum(10)
	ret3 := getSum(10, 20)
	ret4 := getSum(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4)
}
