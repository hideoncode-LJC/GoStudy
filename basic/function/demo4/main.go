package main

import "fmt"

/*
1.全局变量
2.局部变量
*/

var x int = 10

func function() {
	var x int = 100
	fmt.Println("局部变量", x)
}

func main() {
	function()
}
