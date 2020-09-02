package controllers

import (
	"pokedex-web-app/helpers"

	"github.com/astaxie/beego"
)

// PokemonController ...
type PokemonController struct {
	beego.Controller
}

// GetPokemon ...
// @router /pokemon [get]
func (p *PokemonController) GetPokemon() {
	nom := p.GetString("nombre")
	pokemon := helpers.ObtenerPokemon(nom)

	p.Data["json"] = pokemon
	p.ServeJSON()
}

// GetGeneracion ...
func (p *PokemonController) GetGeneracion() {
	id := p.GetString("id")

	gen := helpers.ObtenerGeneracion(id)

	p.Data["json"] = gen
	p.ServeJSON()
}
