package routers

import (
	"github.com/goweb3/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/guest/register", &controllers.UserController{}, "post:Register")
}