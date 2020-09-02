package controllers

import (
	"pokedex-web-app/helpers"

	"github.com/astaxie/beego"
)

// PokemonPreviewController ...
type PokemonPreviewController struct {
	beego.Controller
}

// GetAllPokemons ...
// @router /pokemons [get]
func (p *PokemonPreviewController) GetAllPokemons() {
	pokemons := helpers.GetAllPokemons()
	p.Data["json"] = pokemons
	p.ServeJSON()
}
