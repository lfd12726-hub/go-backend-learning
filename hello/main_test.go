package main

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"普通", "Go", "你好，Go！欢迎学习 Go 后端开发。"},
		{"空", "", "你好，！欢迎学习 Go 后端开发。"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.in); got != tt.want {
				t.Errorf("Greet(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
