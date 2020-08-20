package helpers

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go"
)

func ObtenerPokemon(nombre string) map[string]string {
	prueba := make(map[string]string)
	pok, err := pokeapi.Pokemon(nombre)
	fmt.Println("Nombre:", pok.Name)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	prueba[nombre] = pok.Name
	return prueba
}
