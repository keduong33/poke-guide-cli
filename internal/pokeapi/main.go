package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// some fields are omitted
type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokemon struct {
	Name  string        `json:"name"`
	Moves []PokemonMove `json:"moves"`
}

type PokemonMove struct {
	Move NamedAPIResource `json:"move"`
}

type PokemonDetailedMove struct {
	Name     string `json:"name"`
	Accuracy int    `json:"accuracy"`
}

func GetPokemon(input string) (Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", input)

	response, err := http.Get(url)
	if err != nil {
		return Pokemon{}, errors.New("failed to fetch data from the URL")
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		return Pokemon{}, errors.New("failed to fetch Pokemon: " + fmt.Sprint(response.StatusCode))
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, errors.New("failed to read the response body")
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, errors.New("failed to unmarshal JSON")
	}

	fmt.Println("Moves:")
	for _, move := range pokemon.Moves {
		fmt.Printf("- %s\n", move.Move.Name)
	}

	return pokemon, nil
}
