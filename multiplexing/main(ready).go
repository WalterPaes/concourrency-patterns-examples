package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	canal := multiplexing(printGenerator("Olá Mundo!"), printGenerator("Programando em Go!"))

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(<-canal)
// 	}
// }

// func multiplexing(ch, ch2 <-chan string) chan string {
// 	newCh := make(chan string)

// 	go func() {
// 		for {
// 			select {
// 			case msg := <-ch:
// 				newCh <- msg
// 			case msg := <-ch2:
// 				newCh <- msg
// 			}
// 		}
// 	}()
// 	return newCh
// }

// func printGenerator(word string) <-chan string {
// 	ch := make(chan string)

// 	go func() {
// 		for {
// 			ch <- word
// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}()

// 	return ch
// }
