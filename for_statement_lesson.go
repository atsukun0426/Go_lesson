package main

import "fmt"

func main() {
	/*
		s := []int{}
		for i := 0; i < 5; i++ {
			s = append(s, i)
		}
		fmt.Println(s)
	*/

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			// continue以下のコードが処理されずforに戻る
			continue
		}

		if i > 5 {
			fmt.Println("break")
			// breakは強制的にループを抜ける
			break
		}
		fmt.Println(i)
	}

	// 変数をfor文の外で定義する場合
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
