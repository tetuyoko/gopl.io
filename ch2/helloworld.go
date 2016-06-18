package main

import "fmt"

func HelloWord() string {
	val := "Hello, 世界"
	return val
}

func main() {
	val := HelloWord()
	fmt.Println(val)
}
