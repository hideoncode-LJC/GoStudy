# 变量
## 标识符
+ 只能由字母、下划线、数字组成
+ 只能由字母、下划线开头
## 关键字
+ break，跳出循环
+ default，switch默认值
+ func，声明函数关键字
+ interface，声明接口关键字
+ select
+ case，switch中的每个分支
+ defer
+ go，创建携程
+ map，字典关键字
+ struct，结构体关键字
+ chan，协程关键字
+ if-else，选择结构
+ error，错误信息接口


## 变量和内存的关系？？？编译原理？
## 变量声明和赋值
### 声明与赋值分开
+ 标准声明
```go
var name string//default ""
name = "123"
var age int//default 0
age = 22
var isOK bool//default false
isOK = true
```
+ 批量声明
```go
var (
	a string
	b int
	c bool
	d float32
)
```
+ int默认为0
+ float默认为0.0
+ string默认为空字符串
+ 布尔变量默认为false
+ 切片、函数、指针的默认为nil
### 声明的同时赋值
```go
var a int = 1
```
+ 不说明变量的类型，进行变量推导
```go
var a = 1
var name, age = "111", 20
```
+ 短变量声明
```go
a := 1
```
> 注意参数的作用域
> 
> := 不能使用在函数外
### 匿名变量
+ 不会给匿名变量分配内存
```go
func get() (int, string) {
	return 10, "111"
}

func main() {
	x, _ := get()
	-, y := get()
}
```
# 常量
+ 声明一个
```go
const PI = 3.1315
const e = 2.7182
```
+ 批量声明
```go
const (
	pi = 3.1415
	e = 2.7182
)
//如果声明多个相同的常量
const (
	n1 = 100
	n2
	n3
)
```
+ iota
```go
const (
	n1 = iota//0
	n2//1
	n3//2
	n4//3
)
//存在省略值
const (
	n1 = iota//0
	n2//1
	_
	n3//3
)
const n5 = iota //0

const (
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
	EB = 1 << (10 * iota)
	ZB = 1 << (10 * iota)
)

//定义多个iota
const (
	n1, n2 = iota + 1, iota + 2
	n3, n4
	n5, n6
)
```
