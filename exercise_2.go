package main

import "fmt"

func getMinNumBySliceExercise() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	var min int
	for i, num := range l {
		if i == 0 {
			min = num
			continue
		}

		if min >= num {
			min = num
		}
	}

	fmt.Println(min)
}

func getMapSum() {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	var sum int
	for _, v := range m {
		sum += v
	}

	fmt.Println(sum)
}

func main() {
	getMinNumBySliceExercise()
	getMapSum()
}
