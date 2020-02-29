// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/customer",
				beego.NSInclude(
					&controllers.MainController{},
				),
			),
		)
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.MainController{}, "get:Get")
}

//
//func init() {
//    beego.Router("/", &controllers.MainController{})
//
//}
