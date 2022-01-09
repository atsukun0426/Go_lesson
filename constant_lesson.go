package main

import "fmt"

//定数はグローバルに宣言することが多い
const Pi = 3.14

const (
	Username = "user"
	Password = "pass"
)

//本来はoverfolewするが定数宣言はuntypeなので実行時に判定される
const big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1)
}
