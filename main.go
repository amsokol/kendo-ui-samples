package main

import (
	"flag"
	"github.com/amsokol/kendo-ui-samples/errors"
	"github.com/amsokol/kendo-ui-samples/models"
	"github.com/amsokol/kendo-ui-samples/models/testdata"
	_ "github.com/amsokol/kendo-ui-samples/routers"
	"github.com/astaxie/beego"
)

var dbHost = flag.String("host", "", "PostgreSQL database host")
var dbName = flag.String("db", "", "PostgreSQL database name")
var dbUser = flag.String("user", "", "PostgreSQL database user")
var dbPwd = flag.String("pwd", "", "PostgreSQL database password")

func main() {
	var err error

	flag.Parse() // Scan the arguments list

	models.DBConn, err = models.InitDB(*dbHost, *dbName, *dbUser, *dbPwd)
	errors.Panic(err)
	defer models.DBConn.Close()

	testdata.CreateTestDate(models.DBConn)

	beego.Run()
}
