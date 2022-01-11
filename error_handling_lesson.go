package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./lesson.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error!!")
	}
	fmt.Println(count, string(data))

	/*下記コードは29行目のように一行で書くことができる
	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("Error")
	}
	*/

	// 変数を他の関数内で使わない場合はこのように書いてもいい
	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}
}
