# fmt库
## Print函数分析
### 调用Print函数
+ 传入空接口参数，即任何类型都可
+ ...表示参数的数目可变
```bash
func Print(...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...), nil
}
```
+ 调用Fprint函数
```go

```