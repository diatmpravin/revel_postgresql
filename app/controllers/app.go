package controllers

import(
  "revel_postgresql/app/models"
  "github.com/revel/revel"
  "fmt"
) 

type App struct {
	ModelController
}

func (c App) New() revel.Result {
  greeting := "Hello, Revel.!!"
	return c.Render(greeting)
}

func (c App) Index(name string) revel.Result {
  fmt.Println("Name------->", name)
  user := models.User{Name: name}
  fmt.Println("Users------->", user)
  fmt.Println("Saving data-->", c.Orm.Save(&user))
  
  fmt.Println("All Data:", c.Orm.Exec("select * from users;"))
  rows, _ := c.Orm.Model(models.User{}).Select("name").Rows()
  fmt.Println("All Data:", rows)
  
  users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Name)
		fmt.Println(err)
		fmt.Println(user)
		users = append(users, user)
	}
	fmt.Println("Users:", users)
	
	return c.Render(users)
}

