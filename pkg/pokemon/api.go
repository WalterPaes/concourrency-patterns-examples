package pokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

func GetByName(name string) string {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var p Pokemon

	//fmt.Println(string(body))

	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("ID: %d | Name: %s | Altura: %d | Peso: %d", p.ID, p.Name, p.Height, p.Weight)
}
