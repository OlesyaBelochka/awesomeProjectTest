package basics_workOutTasks

import (
	"fmt"
)

func Task11() {
	var a, rez uint8

	fmt.Scan(&a)

	//a=0

	if a > 20 {
		rez = a % 10
	} else {
		rez = a
	}
	switch rez {
	case 1:
		fmt.Println(a, "korova")
	case 2, 3, 4:
		fmt.Println(a, "korovy")
	default:
		fmt.Println(a, "korov")
	}
}
