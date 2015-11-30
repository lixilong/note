package controllers
import (
	"github.com/astaxie/beego"
	//"html/template"
	//"net/http"
)
type NoteController struct {
	beego.Controller
}

func (this *NoteController) Get() {
	this.TplNames = "index.tpl"
}