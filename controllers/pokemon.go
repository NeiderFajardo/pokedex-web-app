package controllers

import (
	"pokedex-web-app/helpers"

	"github.com/astaxie/beego"
)

type PokemonController struct {
	beego.Controller
}

// @router /allpokemons/:nombre
func (p *PokemonController) Get() {
	//nombre := p.GetString(":nombre")

	pokemon := helpers.ObtenerPokemon("umbreon")

	p.Data["json"] = pokemon
	p.ServeJSON()
}
