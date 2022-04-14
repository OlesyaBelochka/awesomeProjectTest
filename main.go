package main

//
//import sectionThree_refactoringType "awesomeProject/sectionThree/refactoringTypes"
//
//func main() {
	//basics_workOutTasks.Task9()
	//basics_workOutTasks.Task10()
	//basics_workOutTasks.Task11()
	//sectionThree_parallelTwo.Test1()
	sectionThree_refactoringType.Task3()

//}
/*package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {

	chan_out := make(chan int)

	go func(chan_out chan int) {
		var x int

		select { // Оператор select

		case x = <-firstChan: // Ждет, когда проснется гофер

			chan_out <- x * x

		case x = <-secondChan: // Ждет окончания времени

			chan_out <- x * 3

		default:

		}
	}(chan_out)

	return chan_out

}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})

	r := calculator(ch1, ch2, stop)
	ch1 <- 3
	//ch2 <- 2

	//close(stop)
	fmt.Println(<-r)

}*/
