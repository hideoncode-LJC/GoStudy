package main

import "fmt"

func get() ([]byte, []byte) {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	res := make([]byte, 3)
	copy(res, raw[0:3])
	return raw[0:3], res
}

func main() {
	data1, data2 := get()
	fmt.Println(data1)
	fmt.Println(len(data1), cap(data1), &data1[0])
	fmt.Println(data2)
	fmt.Println(len(data2), cap(data2), &data2[0])

}
