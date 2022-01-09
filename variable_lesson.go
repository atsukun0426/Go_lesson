package main

import "fmt"

func main() {
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	fmt.Println(i, f64, s, t, f)

	//変数宣言の省略は関数内でしか使えない
	xi := 1
	xf64 := 1.25
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)

}
