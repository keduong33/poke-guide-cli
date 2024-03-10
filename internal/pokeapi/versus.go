package pokeapi

import "fmt"

func Versus(attacker, defender string) error {
	fmt.Println(attacker, defender)

	GetPokemon(attacker)

	return nil
}
