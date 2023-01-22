package main

import "fmt"

/*
长度固定的数组
数组长度也是固定的一部分
数组是一个类型
[5]int 与 [10]int 不一样
&a a的地址
*a a地址对应的值
*/

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var b = [5]int{1, 3, 3, 4, 5}
	//var c = [10]int{9 : 10}
	var d = [5]int{1, 2, 3, 4, 5}
	var e = new([5]int)
	for i := range e {
		e[i] = i + 1
	}
	if a == b {
		fmt.Printf("a == b\n")
	}

	if a == d {
		fmt.Printf("a == d\n")
	}

	//if(a == e) {
	//
	//}
}
