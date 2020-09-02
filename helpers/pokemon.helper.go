package helpers

import (
	"pokedex-web-app/models"
	"strconv"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

const generaciones int = 7

// ObtenerPokemon ...
func ObtenerPokemon(nombre string) models.Pokemon {
	var (
		pokemon models.Pokemon
	)
	ruta := "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/"
	pokemon.Habilidad = make(map[int]string)
	pokemon.Tipo = make(map[int]string)
	pokemon.Stats = make(map[string]int)
	pokemon.Imagen = make(map[string]string)

	pok, err := pokeapi.Pokemon(strings.ToLower(nombre))
	if err != nil {
		panic(err)
	}
	pokemon.Nombre = strings.ToUpper(pok.Name)
	numRuta := getNumeroPokemon(pok.Sprites.FrontDefault, false)

	pokemon.Imagen["FrontDefault"] = ruta + numRuta + ".png"
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

// GetAllPokemons ..
func GetAllPokemons() []models.PokemonPreview {
	var (
		pokemons []models.PokemonPreview
		pokemon  models.PokemonPreview
	)
	frontDefault := "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/"
	for i := 0; i < generaciones; i++ {
		pokex, err := pokeapi.Generation(strconv.Itoa(i + 1))
		if err != nil {
			panic(err)
		}
		for _, j := range pokex.PokemonSpecies {

			pokemon.Nombre = strings.Title(strings.ToLower(j.Name))

			pokemon.Preview = frontDefault + getNumeroPokemon(j.URL, true) + ".png"

			pokemons = append(pokemons, pokemon)
		}
	}
	return pokemons
}

// ObtenerGeneracion ...
func ObtenerGeneracion(id string) map[string]interface{} {
	gens := make(map[string]interface{})

	g, err := pokeapi.Generation(id)
	if err != nil {
		panic(err)
	}
	gens["pokemons"] = g.PokemonSpecies
	return gens
}

func getNumeroPokemon(ruta string, front bool) string {
	numero := ""
	for _, char := range ruta {
		if _, err := strconv.Atoi(string(char)); err == nil {
			numero += string(char)
		}
	}
	if front {
		return numero[1:]
	}
	return numero
}
