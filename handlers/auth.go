package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	email    string
	password string
}

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/auth/login.html"))
	tmpl.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Printf("%s %s", email, password)

	db, err := sql.Open("mysql", "root:secret@(localhost:3306)/go-todo?parseTime=true")

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var user User

	{
		err := db.QueryRow("SELECT email, password FROM users WHERE email = ?", email).Scan(&user.email, &user.password)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Fprintf(w, "%#v", user)
}

func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/auth/register.html"))
	tmpl.Execute(w, nil)
}
