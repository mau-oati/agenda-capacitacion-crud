// @APIVersion 1.0.0
// @Title Agenda Test API
// @Description API prueba para agenda.
package routers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/capacitacion-agenda-crud/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/contacto",
			beego.NSInclude(
				&controllers.ContactoController{},
			),
		),

		beego.NSNamespace("/telefono",
			beego.NSInclude(
				&controllers.TelefonoController{},
			),
		),

		beego.NSNamespace("/correo",
			beego.NSInclude(
				&controllers.CorreoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
