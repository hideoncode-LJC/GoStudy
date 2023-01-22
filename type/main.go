package main

import "fmt"

/*
type的用途
1.创建自定义类型
*/

type myStruct struct {
	a int
	b float64
	c string
}

/*
2.类型别名
*/

type i = int8
type i8 int8

/*
声明结构体
*/

type person struct {
	name string
	city string
	age  int8
}

//同样的字段可以写在同一行
type person1 struct {
	name, city string
	age        int8
}

//结构体实例化
func main() {
	var p person
	p.name = "qwe"
	p.city = "北京"
	p.age = 20
	fmt.Printf("p = %v\n", p)
	fmt.Printf("p = %#v\n", p)
}
