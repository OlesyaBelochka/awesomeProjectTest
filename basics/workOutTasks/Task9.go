package basics_workOutTasks

import (
	"fmt"
)

func getDigit(a, sum int64) int64 {
	if a < 10 {
		sum += a
		return sum
	} else {
		sum += a % 10
		//		fmt.Println(a%10, " ", sum)
		return getDigit(a/10, sum)
	}
}

func getArr(a int64) int64 {

	var sum int64

	sum = getDigit(a, sum)

	//fmt.Println(sum)

	if sum < 10 {
		return sum
	}
	return getArr(sum)
}
func Task9() {
	var a, sum int64

	fmt.Scan(&a)

	sum = getArr(a)

	fmt.Println(sum)

}
