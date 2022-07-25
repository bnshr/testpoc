package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func server1(ch chan string) {
	time.Sleep(6 * time.Second)

	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server 2"
}

func main_select() {
	o1 := make(chan string)
	o2 := make(chan string)

	go server1(o1)
	go server2(o2)

	select {
	case s1 := <-o1:
		fmt.Println(s1)
	case s2 := <-o2:
		fmt.Println(s2)
	}

	ch := make(chan string)
	go process(ch)

	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received val", v)
			return
		default:
			fmt.Println("no value")
		}
	}

}
