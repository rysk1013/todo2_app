package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

type Session struct {
	ID int
	UUID string
	Email string
	UserID int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (uuid, name, email, password, created_at) 						VALUES (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}

	cmd := `SELECT id, uuid, name, email, password, created_at 
					FROM users
					WHERE id = ?`

	err = Db.QueryRow(cmd, id).Scan(
			&user.ID,
			&user.UUID,
			&user.Name,
			&user.Email,
			&user.PassWord,
			&user.CreatedAt,
	)

	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `UPDATE users 
					SET name = ?, email = ? 
					WHERE id = ?`

	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `DELETE FROM users
					WHERE id = ?`

	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}

	cmd := `SELECT id, uuid, name, email, password, created_at
					FROM users
					WHERE email = ?`

	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}

	// Create Session
	cmd1 := `INSERT INTO sessions (uuid, email, user_id, created_at)
					 VALUES (?, ?, ?, ?)`

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	// Get Session
	cmd2 := `SELECT id, uuid, email, user_id, created_at
					 FROM sessions
					 WHERE user_id = ? AND email = ?`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
	)

	return session, err
}