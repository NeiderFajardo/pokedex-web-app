package routers

import (
	"pokedex-web-app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/pokemon", &controllers.PokemonController{}, "get:GetPokemon")
	beego.Router("/pokemon/gen", &controllers.PokemonController{}, "get:GetGeneracion")

}
