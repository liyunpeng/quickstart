package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["quickstart/controllers:MainController"] = append(beego.GlobalControllerRouter["quickstart/controllers:MainController"],
        beego.ControllerComments{
            Method: "JudgeLogin",
            Router: `/backlogin/judge`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
