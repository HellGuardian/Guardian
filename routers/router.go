package Guardian

import (
	"github.com/astaxie/beego"
	"Guardian/controllers"
)

func init () {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{})
	//beego.Router("/login", &controllers.LoginController{})
	//beego.Router("/host", &controllers.HostController{})
}