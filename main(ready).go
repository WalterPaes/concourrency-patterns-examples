package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch := make(chan string)

// 	go print("Olá, mundo!", ch)
// 	go print("Hello, World", ch)

// 	for i := 0; i < 5; i++ {
// 		r := <-ch
// 		fmt.Println(r)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func print(word string, ch chan string) {
// 	for {
// 		ch <- fmt.Sprintf("Print: %s", word)
// 	}
// }
