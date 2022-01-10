package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	// 新しい要素を追加する場合
	m["new"] = 500
	fmt.Println(m)

	// 存在しない要素を指定すると0が返る
	// m["xxxx"]の返り値は、値とbool型を返し、bool型は無視できる
	fmt.Println(m["nothing"])

	// 存在の確認
	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 500
	fmt.Println(m2)

	// varでmapやスライスを宣言するとnilになるのでエラーになる
	/*
		var m3 map[string]int
		m3["pc"] = 5000
		fmt.Println(m3)
	*/

	var s []int
	if s == nil {
		fmt.Println("Nil")
	}

}
