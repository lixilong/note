package routers

import (
	"note/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/note", &controllers.NoteController{})
}
