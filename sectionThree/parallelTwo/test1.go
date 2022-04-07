package sectionThree_parallelTwo

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	fmt.Println("сработала функция fibonacci")
	for {
		fmt.Println("сработал цикл в fibonacci")
		select {
		case c <- x:
			fmt.Println("#2 - сработала case c <- x")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func Test1() {
	c := make(chan int)

	quit := make(chan int)
	go func() {
		fmt.Println("сработал кусок go func")
		for i := 0; i < 10; i++ {
			fmt.Println("'это цикл", <-c)

		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
