package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:CorreoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:CorreoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:CorreoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:CorreoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:CorreoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:CorreoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:CorreoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:CorreoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:CorreoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:CorreoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:TelefonoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:TelefonoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:TelefonoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:TelefonoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/capacitacion-agenda-crud/controllers:TelefonoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
