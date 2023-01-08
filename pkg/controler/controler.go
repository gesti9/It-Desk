package controler

import (
	"log"
	"net/http"
	"text/template"
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

	tmpl, err := template.ParseFiles("../html/sign.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, user)

}
func Welcome(w http.ResponseWriter, r *http.Request) {
	user := AuthNaP{}

	tmpl, err := template.ParseFiles("../html/welcome.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, user)

}
func Home(w http.ResponseWriter, r *http.Request) {
	user := AuthNaP{}

	tmpl, err := template.ParseFiles("../html/home.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, user)

}
