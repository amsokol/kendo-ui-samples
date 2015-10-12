package controllers

import (
//	"github.com/amsokol/kendo-ui-samples/kendoui"
	"github.com/astaxie/beego"
	"github.com/amsokol/kendo-ui-samples/models"
	"encoding/json"
)

type AutocompleteController struct {
	beego.Controller
}

// @router /autocomplete [get]
func (c *AutocompleteController) Get() {
	c.TplNames = "autocomplete.html"
	c.Render()
}

// @router /autocomplete/products [get]
func (c *AutocompleteController) GetProducts() {
	products := make([]models.Product, 0)

	//	params := kendoui.ParseParams(c.Input())

	s := c.GetString("filter[filters][0][value]")

	rows, err := models.DBConn.Query("SELECT product_jdoc AS product_name " +
		"FROM products " +
		"WHERE product_jdoc->>'name' ILIKE '%' || $1 || '%'", s)
	if err != nil {
		beego.Error(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var p string
			var product models.Product
			if err := rows.Scan(&p); err != nil {
				beego.Error(err)
			} else if err = json.Unmarshal([]byte(p), &product); err != nil {
				beego.Error(err)
			} else {
				products = append(products, product)
			}
		}
	}

	c.Data["jsonp"] = &products
	c.ServeJsonp()
}
