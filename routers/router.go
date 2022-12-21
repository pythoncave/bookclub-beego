package routers

import (
	"bookclub/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.MainController{}, "get:About")
}
