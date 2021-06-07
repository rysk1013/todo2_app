package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	User_id   int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `INSERT INTO todos(
				content,
				user_id,
				created_at) values(?, ?, ?)`
	
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	
	return err
}