package pokeapi

import (
	"maps"
)

type ConsolidatedDamageGuide struct {
	MovesAgainstDefender []Move
	MovesAgainstAttacker []Move
}

type TypeSet map[string]Type
type MoveSet map[string]Move

func findCommonMoves(allMoves MoveSet, attackerMoves []Move) []Move {
	validMoves := make([]Move, 0)

	for _, move := range attackerMoves {
		if allMoves[move.Name] != (Move{}) {
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
}

func GetAllMovesOfType(pokemonType Type) (MoveSet, error) {
	detailType, err := GetTypeDetails(pokemonType)

	if err != nil {
		return MoveSet{}, err
	}

	allMoves := make(MoveSet)
	for _, move := range detailType.Moves {
		allMoves[move.Name] = Move(move)
	}

	return allMoves, nil
}

func GetWeakAgainstTypes(pokemonType Type) (TypeSet, error) {
	detailType, err := GetTypeDetails(pokemonType)

	if err != nil {
		return TypeSet{}, err
	}

	weakAgainstTypes := make(TypeSet)
	for _, pokemonType := range detailType.DamageRelations.DoubleDamageFrom {
		weakAgainstTypes[pokemonType.Name] = Type(pokemonType)
	}

	return weakAgainstTypes, nil
}

func getEffectiveMovesAgainst(pokemonTypes []Type) (MoveSet, error) {
	weakAgainstTypes := make(TypeSet)
	for _, pokemonType := range pokemonTypes {
		weakTypes, err := GetWeakAgainstTypes(pokemonType)

		if err != nil {
			return MoveSet{}, err
		}
		maps.Copy(weakAgainstTypes, weakTypes)
	}

	effectiveMoves := make(MoveSet)
	for _, weakAgainstType := range weakAgainstTypes {
		moves, err := GetAllMovesOfType(weakAgainstType)

		if err != nil {
			return MoveSet{}, err
		}

		maps.Copy(effectiveMoves, moves)
	}

	return effectiveMoves, nil
}

func (attacker Pokemon) getMovesAgainstThePokemon(defender Pokemon) ([]Move, error) {
	allMovesAgainstDefender, err := getEffectiveMovesAgainst(defender.Types)
	if err != nil {
		return nil, err
	}

	movesAgainstDefender := findCommonMoves(allMovesAgainstDefender, attacker.Moves)

	return movesAgainstDefender, nil
}

func GetDamageGuide(attacker, defender Pokemon) (ConsolidatedDamageGuide, error) {

	movesAgainstDefender, err := attacker.getMovesAgainstThePokemon(defender)

	if err != nil {
		return ConsolidatedDamageGuide{}, nil
	}

	movesAgainstAttacker, err := defender.getMovesAgainstThePokemon(attacker)

	if err != nil {
		return ConsolidatedDamageGuide{}, nil
	}

	return ConsolidatedDamageGuide{movesAgainstDefender, movesAgainstAttacker}, nil
}

func Versus(attackerPokemon, defenderPokemon string) (ConsolidatedDamageGuide, error) {
	attacker, err := GetPokemon(attackerPokemon)

	if err != nil {
		return ConsolidatedDamageGuide{}, err
	}

	defender, err := GetPokemon(defenderPokemon)

	if err != nil {
		return ConsolidatedDamageGuide{}, err
	}

	damageGuide, err := GetDamageGuide(attacker, defender)

	if err != nil {
		return ConsolidatedDamageGuide{}, err
	}

	return damageGuide, nil
}
