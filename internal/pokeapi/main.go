package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"pokedex-cli/internal/utils"
	"pokedex-cli/internal/utils/myAxios"
)

var GameVersionGroup = "red-blue"

type Pokemon struct {
	Name  string
	Moves []Move
	Types []Type
}

type SimplifiedInfo struct {
	Name string
	Url  string
}
type Move SimplifiedInfo
type Type SimplifiedInfo

func validatePokemonName(name string) (string, error) {
	if name == "" {
		return "", errors.New("pokemon's name cannot be empty")
	}
	return name, nil
}

func GetPokemon(pokemonName string) (Pokemon, error) {
	validatedName, err := validatePokemonName(pokemonName)

	if err != nil {
		return Pokemon{}, err
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", validatedName)

	data, err := myAxios.GetRequest(url)

	if err != nil {
		return Pokemon{}, err
	}

	var apiPokemon ApiPokemon
	err = json.Unmarshal(data, &apiPokemon)
	if err != nil {
		return Pokemon{}, utils.CreateJsonReadError(err)
	}

	pokemon := convertApiPokemonToPokemon(apiPokemon)

	return pokemon, nil
}

func GetTypeDetails(pokemonType Type) (ApiPokemonDetailedType, error) {
	url := pokemonType.Url
	data, err := myAxios.GetRequest(url)

	var detailType ApiPokemonDetailedType

	if err != nil {
		return ApiPokemonDetailedType{}, nil
	}

	if err := json.Unmarshal(data, &detailType); err != nil {
		return ApiPokemonDetailedType{}, utils.CreateJsonReadError(err)
	}

	return detailType, nil
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
