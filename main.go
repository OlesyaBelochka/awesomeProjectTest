package main

import (
	"fmt"
)

func getArr(a int) ([]int, int) {
	var arr []int
	var countA int
	if a > 10 {
		countA = 1
	}
	if a >= 10 && a < 100 {
		countA = 2
	}
	if a >= 100 && a < 1000 {
		countA = 3
	}
	if a >= 1000 && a < 10000 {
		countA = 4
	}
	if a == 10000 {
		countA = 5
	}

	for i := 0; i < countA; i++ {

		arr = append(arr, a%10)
		a = a / 10
	}

	return arr, countA
}
func main() {
	var a, b int
	var arr, brr []int
	var countA, countB int

	fmt.Scan(&a, &b)

	arr, countA = getArr(a)
	brr, countB = getArr(b)

	for i := 0; i < countA; i++ {

		for j := 0; j < countB; j++ {
			if arr[i] == brr[j] {
				defer fmt.Print(arr[i], " ")
			}
		}

	}

}
