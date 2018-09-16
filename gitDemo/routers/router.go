package routers

import (
	"github.com/astaxie/beego"
	"gittest/gitDemo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
