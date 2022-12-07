package controler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"workspace/service"
)

type AuthNaP struct {
	Login string
	Pass  string
}

type Task struct {
	Name string
	Task string
}

func Auth(w http.ResponseWriter, r *http.Request) {
	user := AuthNaP{}

	tmpl, err := template.ParseFiles("html/start.html")
	if err != nil {
		log.Fatal(err)
	}

	switch r.Method {
	case "GET":
		tmpl.Execute(w, user)
	case "POST":
		if service.UserList(r.FormValue("login"), r.FormValue("pass")) {
			tmpl, _ := template.ParseFiles("html/home.html")
			tmpl.Execute(w, user)
		} else {
			fmt.Println("wrong")
			tmpl.Execute(w, user)
		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	text := Task{}
	tmpl, err := template.ParseFiles("html/home.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, text)
}
