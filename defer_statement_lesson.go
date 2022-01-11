package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("Hello foo")
	fmt.Println("World")
}

func main() {
	/*
		defer fmt.Println("World foo")
		fmt.Println("Hello")
		foo()
	*/
	// file.Open実行後deferですぐにfile.Closeを書くことで書き忘れがなくなる
	file, _ := os.Open("./lesson.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
