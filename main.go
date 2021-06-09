package main

import (
	"fmt"
	// "log"
	// "todo2_app/config"
	"todo2_app/app/models"
)

func main() {
	fmt.Println(models.Db)
	
	// u := &models.User{} 
	// u.Name = "test"
	// u.Email = "text@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)
	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)
	// u.Name = "Test2"
	// u.Email = "test2@example.com"

	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	user, _ := models.GetUser(2)
	user.CreateTodo("Second Todo")

	todos, _ := models.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	}
}
