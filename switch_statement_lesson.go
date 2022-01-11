package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {
	os := getOsName()
	// defaultは書かなくてもいい
	switch os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("default!!")
	}

	// switch文の中で変数を定義することもできるが変数をswitch文の外から呼び出すことは出来ない
	switch os2 := getOsName(); os2 {
	case "mac":
		fmt.Println("Mac!!", os)
	case "windows":
		fmt.Println("Windows!!", os)
	default:
		fmt.Println("default!!", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() > 17:
		fmt.Println("Afternoon")
	}
}
