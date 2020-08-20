package routers

import (
	"pokedex-web-app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/allpokemons", &controllers.PokemonController{})
}
