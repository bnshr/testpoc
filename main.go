package main

import (
	"fmt"
	"time"
)

/*
func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {

	go numbers()
	go alphabets()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println()
	fmt.Println("Main terminated")

	var a chan int

	if a == nil {
		fmt.Println("channel a is nil, going to define it")

		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}

	data := <-a
	a <- data
}*/

func hello(done chan bool) {
	fmt.Println("going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("awake and going to write to done")
	done <- true
}

func calcSquares(number int, op chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	op <- sum
}

func calcCubes(number int, op chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	op <- sum
}

func main() {
	/*done := make(chan bool)
	fmt.Println("main going to call hello")
	go hello(done)
	<-done
	fmt.Println("main function")*/

	number := 999
	sqr := make(chan int)
	cube := make(chan int)

	go calcSquares(number, sqr)
	go calcCubes(number, cube)

	squares, cubes := <-sqr, <-cube

	fmt.Println("Final output", squares, cubes)
}
