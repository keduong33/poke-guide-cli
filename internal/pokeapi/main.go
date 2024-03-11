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
	Types []Type
}

type Move struct {
	Name string
	Url  string
}

type Type struct {
	Name string
	Url  string
}

func GetPokemon(pokemonName string) (Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)

	data, err := myAxios.GetRequest(url)

	if err != nil {
		return Pokemon{}, err
	}

	var apiPokemon ApiPokemon
	err = json.Unmarshal(data, &apiPokemon)
	if err != nil {
		return Pokemon{}, errors.New("failed to understand JSON - " + err.Error())
	}

	pokemon := convertApiPokemonToPokemon(apiPokemon)

	return pokemon, nil
}

func convertApiPokemonToPokemon(apiPokemon ApiPokemon) Pokemon {
	var moves []Move
	for _, move := range apiPokemon.Moves {
		for _, versionGroupDetail := range move.VersionGroupDetails {
			if versionGroupDetail.VersionGroup.Name == GameVersionGroup {
				moves = append(moves, Move{Name: move.Move.Name, Url: move.Move.Url})
				continue
			}
		}
	}

	var types []Type

	for _, pokemonType := range apiPokemon.Types {
		types = append(types, Type{Name: pokemonType.Type.Name, Url: pokemonType.Type.Url})
	}

	return Pokemon{
		Name:  apiPokemon.Name,
		Moves: moves,
		Types: types,
	}
}
