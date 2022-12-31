package main

import (
	"fmt"
	"todo-app/app/controllers"
	"todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

	// u := &models.User{}
	// u.Name = "hashitaro"
	// u.Email = "hashitaro@example.com"
	// u.Password = "hashitaro"
	// fmt.Println(u)
	// u.CreateUser()

	// user, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
}
