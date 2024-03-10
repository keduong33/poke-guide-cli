package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"pokedex-cli/internal/utils/myAxios"
)

var GameVersionGroup = "red-blue"

type Pokemon struct {
	Name  string
	Moves []Move
}

type Move struct {
	Name string
	Url  string
}

func GetPokemonMoves(input string) ([]Move, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", input)

	data, err := myAxios.GetRequest(url)

	if err != nil {
		return nil, err
	}

	var pokemon ApiPokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return nil, errors.New("failed to unmarshal JSON - " + err.Error())
	}

	var moves []Move

	for _, move := range pokemon.Moves {
		for _, versionGroupDetail := range move.VersionGroupDetails {
			if versionGroupDetail.VersionGroup.Name == GameVersionGroup {
				moves = append(moves, Move{Name: move.Move.Name, Url: move.Move.Url})
				continue
			}
		}
	}

	return moves, nil
}
