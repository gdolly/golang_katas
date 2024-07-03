package main

import "fmt"

func sum(array_numbers []int, channel1 chan int) {
	sum := 0
	for _, v := range array_numbers {
		sum += v
	}
	channel1 <- sum // send sum to channel1

	channel1 <- 10
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	channel1 := make(chan int)
	println(channel1)

	//channel1 <- 10

	//af := <-channel1
	//println(af)

	go sum(s[:len(s)/2], channel1)

	go sum(s[len(s)/2:], channel1)

	x, y, z := <-channel1, <-channel1, <-channel1 // receive from channel1

	fmt.Println(x, y, x+y, z)
}
