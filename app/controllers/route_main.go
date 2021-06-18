package controllers

import (
	"log"
	"net/http"
	"todo2_app/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}

		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Fatalln(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}

		todo, err := models.GetTodo(id)
		if err != nil {
			log.Fatalln(err)
		}

		generateHTML(w, todo, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}

		content := r.PostFormValue("content")
		
		todo := &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := todo.UpdateTodo(); err != nil {
			log.Fatalln(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}

		todo, err := models.GetTodo(id)
		if err != nil {
			log.Fatalln(err)
		}

		if err := todo.DeleteTodo(); err != nil {
			log.Fatalln(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

func userEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}

		generateHTML(w, user, "layout", "private_navbar", "user_edit")
	}
}

func userUpdate(w http.ResponseWriter, r *http.Request, id int) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		name := r.PostFormValue("name")
		email := r.PostFormValue("email")

		user := &models.User{ID: id, Name: name, Email: email}
		if err := user.UpdateUser(); err != nil {
			log.Fatalln(err)
		}

		if err := user.UpdateSession(); err != nil {
			log.Fatalln(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

func userDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}

		if err := user.DeleteTodos(); err != nil {
			log.Fatalln(err)
		}

		if err := sess.DeleteSessionByUUID(); err != nil {
			log.Fatalln(err)
		}

		if err := user.DeleteUser(); err != nil {
			log.Fatalln(err)
		}

		http.Redirect(w, r, "/signup", 302)
	}
}