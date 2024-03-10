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
	Move                NamedAPIResource     `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonDetailedMove struct {
	Name     string `json:"name"`
	Accuracy int    `json:"accuracy"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	LevelLearnedAt  int              `json:"level_learned_at"`
}

var GameVersionGroup = "red-blue"

type AvailableMove = NamedAPIResource

func GetPokemonMoves(input string) ([]AvailableMove, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", input)

	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("failed to fetch data from the URL")
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		return nil, errors.New("failed to fetch Pokemon: " + fmt.Sprint(response.StatusCode))
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("failed to read the response body - " + err.Error())
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return nil, errors.New("failed to unmarshal JSON - " + err.Error())
	}

	var moves []AvailableMove

	for _, move := range pokemon.Moves {
		for _, versionGroupDetail := range move.VersionGroupDetails {
			if versionGroupDetail.VersionGroup.Name == GameVersionGroup {
				moves = append(moves, move.Move)
				continue
			}
		}
	}

	return moves, nil
}
