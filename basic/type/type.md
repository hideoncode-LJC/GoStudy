# type
## 基础数据类型
+ 整数
  + int8 = byte
  + int16
  + int32 = rune
  + int64
  + uint8
  + uint16
  + uint32
  + uint64
+ 特殊类型
  + uint，32位机器就是32位，64位机器就是64位
  + int，32位机器就是32位，64位机器就是64位
  + uintptr，指针类型
## 数字字面量语法 Number literals syntax
```go
//声明二进制和输出二进制
a := 0b101010
fmt.Printf("%b\n", a)
//声明和输出八进制
b := 0o1234567 or b:= 01234567
fmt.Printf("%o\n", b)
c := 0x1234
fmt.Printf("%x\n", c)
fmt.Printf("%X\n", c)
```
## 浮点数类型
```go
package main

import (
	"fmt"
	"math"
)
//声明浮点数
var f32 float32 = 1.0
var f64 float64 = 1.123456
//输出浮点数
fmt.Printf("%f\n", f32)
//保留多少位数的浮点数
fmt.Printf("%0.10f\n", f64)
//float32和float64的最大值
fmt.Printf("%f\n", math.MaxFloat32)
fmt.Printf("%0.10f\n", math.MaxFloat64)
```
## 实数和虚数
+ 分为complex64和complex128 (32 + 32) or (64 + 64)
```go
var c1 complex64 = 1 + i
var c2 complex128 = 10 + 10i
fmt.Println(c1, c2)
```
## 布尔值
+ 默认未false
+ 不允许强制转换
+ 无法参与运算
+ 不参与转换
## 字符串
## 转义字符
+ \r reverse 输入go\rstudy 输出 studygo
+ \n 换行
+ \" “
+ \' '
+ \\ \
+ ``