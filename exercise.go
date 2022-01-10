package main

import "fmt"

func main() {
	// Q1 f := 1.11 をint型に変換
	f := 1.111
	i := int(f)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T %v\n", i, i)

	// Q2 以下の出力結果を答えてください
	/*
		s := []int{1, 2, 5, 6, 2, 3, 1}
		fmt.Println(s[2:4])
	*/

	//Q2 A. [5 6]

	//Q3 以下のような出力結果となるmを作成してください
	/*
		map[string]int map[mike:20 Nancy:24 Messi:30]
	*/
	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v", m, m)
}
