package pokeapi

import (
	"encoding/json"
	"fmt"
	"pokedex-cli/internal/utils/myAxios"
)

type DamageGuide struct {
	MoveName    string
	Recommended bool
	Damage      int
	Accuracy    int
}

func GetDetailType(pokemonType Type) (ApiPokemonDetailedType, error) {
	println(pokemonType.Url)
	var detailType ApiPokemonDetailedType

	url := pokemonType.Url
	data, err := myAxios.GetRequest(url)

	if err != nil {
		return ApiPokemonDetailedType{}, nil
	}

	json.Unmarshal(data, &detailType)

	fmt.Printf("%v", detailType)

	return ApiPokemonDetailedType{}, nil
}

func GetDamageGuide(moves []Move, targetTypes []Type) ([]DamageGuide, error) {
	GetDetailType(targetTypes[0])

	// var guide []DamageGuide
	// // for _, move := range moves {
	// url := fmt.Sprintf("https://pokeapi.co/api/v2/move/%s", moves[0])
	// data, err := myAxios.GetRequest(url)

	// if err != nil {
	// 	println(err.Error())
	// 	return nil, err
	// }

	// var apiPokemonDetailMove ApiPokemonDetailMove

	// json.Unmarshal(data, &apiPokemonDetailMove)
	// println(data)

	// // }

	// return guide, nil
	return nil, nil
}

func Versus(attackerPokemon, defenderPokemon string) ([]DamageGuide, error) {
	fmt.Println(attackerPokemon, defenderPokemon)

	attacker, err := GetPokemon(attackerPokemon)

	if err != nil {
		return nil, err
	}

	defender, err := GetPokemon(defenderPokemon)

	if err != nil {
		return nil, err
	}

	attackDamageGuide, err := GetDamageGuide(attacker.Moves, defender.Types)

	if err != nil {
		return nil, err
	}

	return attackDamageGuide, nil

	// defenceDamageGuide, err := GetDamageGuide(defender.Moves, attacker.Types)

	// if err != nil {
	// 	return nil, err
	// }

	// // guide, err := TargetAgainst(attackerMoves, defender)

	// return nil, nil
}
