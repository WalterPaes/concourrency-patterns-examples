package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/WalterPaes/concourrency-patterns-examples/pkg/pokemon"
// )

// func main() {
// 	start := time.Now()

// 	pokemons := []string{
// 		"pikachu",
// 		"bulbasaur",
// 		"charmander",
// 		"squirtle",
// 		"caterpie",
// 		"weedle",
// 		"pidgey",
// 		"rattata",
// 		"poliwag",
// 	}

// 	size := len(pokemons)
// 	jobs := make(chan string, size)
// 	results := make(chan string, size)

// 	go pokemonWorker(jobs, results)

// 	for _, pokemon := range pokemons {
// 		jobs <- pokemon
// 	}
// 	close(jobs)

// 	for i := 0; i < size; i++ {
// 		r := <-results
// 		fmt.Println(i, r)
// 	}

// 	elapsed := time.Since(start)
// 	log.Printf("FINISHED! TIME: %s", elapsed)
// }

// func pokemonWorker(jobs <-chan string, results chan<- string) {
// 	for name := range jobs {
// 		results <- pokemon.GetByName(name)
// 	}
// }
