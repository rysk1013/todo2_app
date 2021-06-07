package main

import (
	"fmt"
	// "log"
	// "todo2_app/config"
	"todo2_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)
	
	u := &models.User{} 
	u.Name = "test"
	u.Email = "text@example.com"
	u.PassWord = "testtest"
	fmt.Println(u)
	u.CreateUser()
}
