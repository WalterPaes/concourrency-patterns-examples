package main

import (
	"fmt"
	"time"
)

func main() {

}

func printGenerator(word string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			ch <- fmt.Sprintf(word)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return ch
}
