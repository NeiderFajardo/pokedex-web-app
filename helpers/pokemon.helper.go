package helpers

import (
	"pokedex-web-app/models"

	"github.com/mtslzr/pokeapi-go"
)

func ObtenerPokemon(nombre string) models.Pokemon {
	var pokemon models.Pokemon
	pokemon.Habilidad = make(map[int]string)
	pokemon.Tipo = make(map[int]string)
	pokemon.Stats = make(map[string]int)

	pok, err := pokeapi.Pokemon(nombre)
	if err != nil {
		panic(err)
	}
	pokemon.Nombre = pok.Name
	for i := 0; i < len(pok.Stats); i++ {
		pokemon.Stats[pok.Stats[i].Stat.Name] = pok.Stats[i].BaseStat
	}
	for i := 0; i < len(pok.Abilities); i++ {
		pokemon.Habilidad[i] = pok.Abilities[i].Ability.Name
	}

	for i := 0; i < len(pok.Types); i++ {
		pokemon.Tipo[i] = pok.Types[i].Type.Name
	}
	return pokemon

}

func ObtenerGeneracion(id string) map[string]interface{} {
	gens := make(map[string]interface{})

	g, err := pokeapi.Generation(id)
	if err != nil {
		panic(err)
	}
	gens["pokemons"] = g.PokemonSpecies
	return gens
}
