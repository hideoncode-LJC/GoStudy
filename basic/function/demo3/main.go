package main

/*
返回值
1.返回值只有一个可以不用括号
2.返回值有多个必须得加括号
3.可以对返回值变量命名
4.返回值补充
*/

func functionOne(x, y int) int {
	return x + y
}

func functionTwo(x, y int, str string) (string, int) {
	sum := x + y
	return str, sum
}

func functionThree(x, y int, str string) (rx int, rs string) {
	rx = x + y
	rs = str
	return rx, rs
}

/*
interface未声明之前是nil
声明之后如果为空就是 []int{}
*/
func functionFour(str string) []int {
	if str == "" {
		return []int{}
	}
	return nil
}

func main() {

}
