package main

import "fmt"

func main() {
	l := []string{"Python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
	}

	// valueだけを取得したい場合はkeyをアンダースコアで書く必要がある
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	//keyだけを取得したい場合はスライスと違いアンダースコアは省略できる
	for k := range m {
		fmt.Println(k)
	}

	// valueだけを取得したい場合はkeyをアンダースコアで書く必要がある
	for _, v := range m {
		fmt.Println(v)
	}
}
