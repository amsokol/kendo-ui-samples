package routers

import (
	"github.com/amsokol/kendo-ui-samples/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/datepicker", &controllers.DatepickerController{})
}
