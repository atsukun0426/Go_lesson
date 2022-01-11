package main

import "fmt"

/*
  基本的にpanicを自分で書くことはない！！
  しっかりとエラーハンドリングをする！！
*/

func thirdPartyConnectDB() {
	// プログラムが強制終了する
	panic("Unable to connect DB!")
}

func save() {
	defer func() {
		// panicをキャッチしてプログラムが強制終了しないようにする
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
}
