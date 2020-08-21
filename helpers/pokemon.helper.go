package helpers

import (
	"github.com/mtslzr/pokeapi-go"
)

func ObtenerPokemon(nombre string) map[string]interface{} {
	prueba := make(map[string]interface{})
	pok, err := pokeapi.Pokemon(nombre)
	if err != nil {
		panic(err)
	}
	prueba["nombre"] = pok.Name
	prueba["orden"] = pok.Order
	prueba["especies"] = pok.Species
	prueba["stats"] = pok.Stats
	return prueba
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
