package controllers

import (
	"fmt"
	"pokedex-web-app/helpers"

	"github.com/astaxie/beego"
)

type PokemonController struct {
	beego.Controller
}

// @router /pokemon [get]
func (p *PokemonController) GetPokemon() {
	nom := p.GetString("nombre")
	fmt.Println("ENTRA:", nom)
	pokemon := helpers.ObtenerPokemon(nom)

	p.Data["json"] = pokemon
	p.ServeJSON()
}

func (p *PokemonController) GetGeneracion() {
	id := p.GetString("id")

	gen := helpers.ObtenerGeneracion(id)

	p.Data["json"] = gen
	p.ServeJSON()
}
