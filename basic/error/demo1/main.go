package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./split")
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	fmt.Printf("%T\n", file)
	fmt.Println(file)
	errors.New()
}
