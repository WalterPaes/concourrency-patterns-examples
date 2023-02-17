package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch := printGenerator("Olá, mundo")
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(<-ch)
// 	}
// }

// func printGenerator(word string) <-chan string {
// 	ch := make(chan string)

// 	go func() {
// 		for {
// 			ch <- fmt.Sprintf(word)
// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}()

// 	return ch
// }
