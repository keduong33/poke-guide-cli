package pokeapi

import "testing"

func TestFindCommonMovesWithEmptyMoveSet(t *testing.T) {
	allMoves := make(map[string]Move)
	attackerMoves := []Move{{Name: "Attack 1"}, {Name: "Attack 2"}}
	validMoves := findCommonMoves(allMoves, attackerMoves)
	if len(validMoves) != 0 {
		t.Fatalf("findCommonMoves(empty moveSet) should return nothing")
	}
}

func TestFindCommonMovesWithEmptyAttackerMoves(t *testing.T) {
	allMoves := map[string]Move{"Attack 1": {Name: "Attack 1"}, "Attack 3": {Name: "Attack 3"}}
	attackerMoves := []Move{}
	validMoves := findCommonMoves(allMoves, attackerMoves)
	if len(validMoves) != 0 {
		t.Fatalf("findCommonMovesWithEmptyAttackMoves should return nothing")
	}
}

func TestFindCommonMovesWithCommon(t *testing.T) {
	allMoves := map[string]Move{"Attack 1": {Name: "Attack 1"}, "Attack 3": {Name: "Attack 3"}}
	attackerMoves := []Move{{Name: "Attack 1"}, {Name: "Attack 2"}}
	validMoves := findCommonMoves(allMoves, attackerMoves)
	if len(validMoves) != 1 {
		t.Fatalf("findCommonMovesWithCommon should return 1 match value")
	}
}

func TestFindCommonMovesWithNoCommon(t *testing.T) {
	allMoves := map[string]Move{"Attack 1": {Name: "Attack 1"}, "Attack 3": {Name: "Attack 3"}}
	attackerMoves := []Move{{Name: "Attack 2"}, {Name: "Attack 4"}}
	validMoves := findCommonMoves(allMoves, attackerMoves)
	if len(validMoves) != 0 {
		t.Fatalf("findCommonMovesWithNoCommon should return nothing")
	}
}
