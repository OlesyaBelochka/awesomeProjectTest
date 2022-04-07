package basics_workOutTasks

import "fmt"

func Task10() {
	var a, b, i, max int32

	fmt.Scan(&a, &b)

	max = -2147483648

	if a <= b {
		for i = b; i > a; i-- {
			if i%7 == 0 {
				max = i
				break
			}
		}

		if max == -2147483648 {
			fmt.Println("NO")
		} else {
			fmt.Println(max)
		}
	}
}
