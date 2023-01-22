# 函数
## 概述
+ 使用func关键字
+ 函数名由字母下划线组成，同变量命名一样，由字母、数字、下划线构成，函数名的第一个字母不能是数字
+ 同一个包内，函数名不能重复
+ 返回值只有一个可以不加括号，有多个必须加括号
+ 参数和返回值可以有可以没有
```bash
func 函数名(参数)(返回值) {
  函数体
}
```
## 变量作用域
### 全局变量
```go
package main

import "fmt"

var num int64 = 10

func test() {
	fmt.Println("输出全部变量", num)
}

func main() {
	test()
}
```
### 局部变量
+ 函数内部定义的变量
```go
package main

import "fmt"

func function() {
	var x int = 45
	//函数内部的变量只能再函数内部访问
	fmt.Println("局部变量x", x)
}

func main() {
    function()
	//fmt.Println(x)
}
```
+ 如果存在名字一样的全部变量和局部变量，优先访问局部变量
```go
package main

import "fmt"

var x int = 1

func function() {
	var x int = 10
	fmt.Println("局部变量x", x)
}
```
+ if、for、switch 内部声明的变量只能在内部使用
## 函数类型与变量
+ 定义函数类型，函数签名相同的函数可以被归为同一个类型
```go
package main
//demo5
import "fmt"
type calculation func(int, int) int

func add(x,y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	//此时c的类型未 calculation
	//如果使用:= c的类型为 func(int, int) int
	var c calculation = add
	ret1 := c(1 ,2)
	c = sub
	ret2 := c(1, 2)
	fmt.Println(ret1, ret2)
}
```
## 高阶导数
+ 函数作参数、返回值 -- 见demo6
## 匿名函数 
+ 见demo7
+ go语言函数内部只能定义匿名函数
+ 匿名函数可以当作返回值
+ 匿名函数被赋值
+ 匿名函数直接调用
## 闭包函数
+ 见demo8
+ 闭包 = 函数 + 引用环境
## defer
+ 见demo9

