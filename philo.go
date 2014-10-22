package main

import "fmt"
import "time"

func main() {

	var left_fork [5]chan bool
	var right_fork [5]chan bool

	for i := range left_fork {
		left_fork[i] = make(chan bool)
		right_fork[i] = make(chan bool)
	}

	fork := func(i int) {
		left := left_fork[i]
		right := right_fork[i]
		for {
			if i%2 == 0 {
				right <- true
				<-right
				left <- true
				<-left
			} else {
				left <- true
				<-left
				right <- true
				<-right

			}

		}
	}

	philo := func(i int) {
		left := left_fork[i]
		right := right_fork[(i+1)%5]
		for {
			<-right
			<-left
			fmt.Println("Eating")
			left <- true
			right <- true
		}
	}

	for i := 0; i < 5; i++ {
		go fork(i)
	}

	for i := 0; i < 5; i++ {
		go philo(i)
	}

	time.Sleep(1000 * time.Millisecond)

}
