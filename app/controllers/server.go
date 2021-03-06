package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"todo2_app/app/models"
	"todo2_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filename ...string) {
	var files []string
	for _, file := range filename {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

// Get ID from URL(todos)
var validPath = regexp.MustCompile("^/todos/(edit|update|delete|show)/([0-9]+)")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		// todos/edit/"id"
		q := validPath.FindStringSubmatch(r.URL.Path)
		if 	q == nil {
			http.NotFound(w, r)
			return
		}

		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
}

// Get ID from URL(users)
var validPath2 = regexp.MustCompile("^/users/(edit|update|delete)/([0-9]+)")

func parseURL2(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request)  {
		// users/edit/"id"
		q := validPath2.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}

		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static", files))

	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	http.HandleFunc("/todos/show/", parseURL(todoShow))
	http.HandleFunc("/users/edit/", parseURL2(userEdit))
	http.HandleFunc("/users/update/", parseURL2(userUpdate))
	http.HandleFunc("/users/delete/", parseURL2(userDelete))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}