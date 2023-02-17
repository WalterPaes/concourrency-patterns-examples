package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	pokemons := []string{
		"pikachu", "bulbasaur", "charmander",
		"squirtle", "caterpie", "weedle", "pidgey",
		"rattata", "poliwag", "charizard", "snorlax",
		"clefairy", "meowth", "psyduck",
		"alakazam", "porygon", "mewtwo", "mew",
	}
	fmt.Println(pokemons)

	duration := time.Since(start)
	log.Printf("FINISHED! TIME: %f", duration.Seconds())
}
