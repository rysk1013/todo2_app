package main

import (
	"fmt"
	"log"
	// "todo2_app/config"
	"todo2_app/app/models"
	//"todo2_app/app/controllers"
)

func main() {
	fmt.Println(models.Db)
	
	// controllers.StartMainServer()

	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(session)
}
