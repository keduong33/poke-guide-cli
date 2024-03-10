package pokeapi

/*
API structs
*/

type ApiNamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ApiPokemon struct {
	Name  string           `json:"name"`
	Moves []ApiPokemonMove `json:"moves"`
}

type ApiPokemonMove struct {
	Move                ApiNamedAPIResource     `json:"move"`
	VersionGroupDetails []ApiPokemonMoveVersion `json:"version_group_details"`
}

type PokemonDetailedMove struct {
	Name     string `json:"name"`
	Accuracy int    `json:"accuracy"`
}

type ApiPokemonMoveVersion struct {
	MoveLearnMethod ApiNamedAPIResource `json:"move_learn_method"`
	VersionGroup    ApiNamedAPIResource `json:"version_group"`
	LevelLearnedAt  int                 `json:"level_learned_at"`
}
