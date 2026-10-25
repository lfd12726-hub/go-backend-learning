package main

import (
	"fmt"
)

// func main() {
// 	// 第 1 课：变量、类型、函数
// 	title := "买牛奶"
// 	done := false

// 	fmt.Println("任务：", title)
// 	fmt.Println("完成了吗：", done)
// }

func describe(id int, title string, done bool) string {
	var status string
	if done {
		status = "已完成"
	} else {
		status = "未完成"
	}
	return fmt.Sprintf("%d.%s[%s]", id, title, status)
}

func main() {
	id := 1
	fmt.Println(describe(id, "  买牛奶", true))
}
