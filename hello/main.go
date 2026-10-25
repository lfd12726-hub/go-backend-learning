package main

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("你好，%s！欢迎学习 Go 后端开发。", name)
}

func main() {
	var shuliang int
	shuliang = 100
	fmt.Print(shuliang)
}
