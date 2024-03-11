package pokeapi

import "fmt"

func Versus(attacker, defender string) error {
	fmt.Println(attacker, defender)

	attackerMoves, err := GetPokemonMoves(attacker)

	return nil
}
