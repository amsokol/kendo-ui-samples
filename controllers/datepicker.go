package controllers

import (
	"github.com/astaxie/beego"
)

type DatepickerController struct {
	beego.Controller
}

func (c *DatepickerController) Get() {
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "datepicker.html"
}
