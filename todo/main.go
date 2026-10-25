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

type Task struct {
	Id    int
	Title string
	Done  bool
}

func describe(t Task) string {
	var status string
	if t.Done {
		status = "已完成"
	} else {
		status = "未完成"
	}
	return fmt.Sprintf("%d. %s [%s]", t.Id, t.Title, status)
}

// 添加新的任务
func addTask(tasks []Task, title string) []Task {
	// 定义新的任务列表，单个的
	newTask := Task{
		Id:    len(tasks) + 1,
		Title: title,
		Done:  false,
	}
	// 追加到 tasks
	tasks = append(tasks, newTask)
	// 返回新的任务
	return tasks
}

func main() {
	task := []Task{
		{Id: 1, Title: "买牛奶", Done: true},
		{Id: 2, Title: "写代码", Done: false},
		{Id: 3, Title: "跑步", Done: false},
	}
	task = addTask(task, "读书")
	task[0].Done = false
	for _, t := range task {
		fmt.Println(describe(t))
	}
}
