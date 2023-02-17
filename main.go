package main

import (
	"fmt"
	"time"
)

func main() {
	print("Olá, mundo!")
}

func print(word string) {
	for {
		fmt.Println(word)
		time.Sleep(time.Millisecond * 500)
	}
}
