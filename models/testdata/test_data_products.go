package testdata

import (
	"encoding/json"
	"github.com/amsokol/kendo-ui-samples/errors"
	"database/sql"
"github.com/amsokol/kendo-ui-samples/models"
)

func insertProducts(db *sql.DB) {
	db.Exec("TRUNCATE products")

	p := models.Product{Id: "1d2cddf1-2ce1-4ca6-7986-079ca329a77c", Name: "iPhone 6"}
	data, err := json.Marshal(p)
	errors.Panic(err)
	_, err = db.Exec("INSERT INTO products(product_jdoc) VALUES($1)", data)
	errors.Panic(err)

	p = models.Product{Id: "0f9e0446-2c60-4173-be7b-57a7415a1cbe", Name: "iPad 2"}
	data, err = json.Marshal(p)
	errors.Panic(err)
	_, err = db.Exec("INSERT INTO products(product_jdoc) VALUES($1)", data)
	errors.Panic(err)

	p = models.Product{Id: "6c656fe5-0cef-4e0b-9343-e792a69d1b30", Name: "Huawei P8"}
	data, err = json.Marshal(p)
	errors.Panic(err)
	_, err = db.Exec("INSERT INTO products(product_jdoc) VALUES($1)", data)
	errors.Panic(err)

	p = models.Product{Id: "04f22c3c-4939-47d2-a1c2-462fb9a19c6c", Name: "Samsung S6"}
	data, err = json.Marshal(p)
	errors.Panic(err)
	_, err = db.Exec("INSERT INTO products(product_jdoc) VALUES($1)", data)
	errors.Panic(err)

	p = models.Product{Id: "502c332f-b945-4bf6-bd97-55c05b65fb42", Name: "Blackberry Pasport"}
	data, err = json.Marshal(p)
	errors.Panic(err)
	_, err = db.Exec("INSERT INTO products(product_jdoc) VALUES($1)", data)
	errors.Panic(err)
}