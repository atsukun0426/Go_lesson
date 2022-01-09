package main

import "fmt"

func main() {
	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)

	fmt.Printf("type=%T value=%v\n", u8, u8)

	// 演算
	x := 1 + 1
	fmt.Println(x)
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("10 - 1 = ", 10-1)
	fmt.Println("10 / 2 = ", 10/2)
	fmt.Println("5 * 2 = ", 5*2)
	fmt.Println("10 / 3 = ", 10/3)
	fmt.Println("10.00 / 3 = ", 10.0/3)
	fmt.Println("10 / 3 = ", 10/3.0)

	i := 0
	fmt.Println(i)
	// 変数に1だけ足す
	i++
	fmt.Println(i)
	// 変数に1だけ引く
	i--
	fmt.Println(i)

	// 数値のシフト
	fmt.Println(1 << 0) // 0001 0001
	fmt.Println(1 << 1) // 0001 0010
	fmt.Println(1 << 2) // 0001 0100
	fmt.Println(1 << 3) // 0001 1000

}
