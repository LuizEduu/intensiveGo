package main

import (
	"fmt"
	"time"
)

func main() {
	/* go process()
	go process()
	process() */

	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}

	}()

	for j := range channel {
		fmt.Println(j)
		time.Sleep(time.Second)
	}

}

/* func process() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
} */
