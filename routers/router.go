package Guardian

import (
	"github.com/astaxie/beego"
	"Guardian/controllers"
)

func init () {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{})
	beego.Router("/demo/welcome", &controllers.WelcomeController{})
	beego.Router("/demo/scan-hot-key", &controllers.HotKeyController{})
	beego.Router("/demo/data-table", &controllers.DataTableController{})
	beego.Router("/demo/noauth_space", &controllers.NoAuthController{})
}