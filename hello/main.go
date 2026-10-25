package main

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("你好，%s！欢迎学习 Go 后端开发。", name)
}

func main() {
	fmt.Println(Greet("世界"))
}
