package controllers

import (
  "revel_postgresql/app/models"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	r "github.com/revel/revel"
)

var DB gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost user=postgres password=ankit1234 dbname=revel_postgresql sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	} else {
		fmt.Println("Successfully connected to DB: %#v", DB)
	}
}

type ModelController struct {
	*r.Controller
	Orm gorm.DB
}

func (c *ModelController) Begin() r.Result {
  fmt.Println("DB HERE---->%+v", DB)
	c.Orm = DB
	return nil
}

func (c *ModelController) AddTables() r.Result {
  c.Orm.CreateTable(models.User{})
  return nil
}



