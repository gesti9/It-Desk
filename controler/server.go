package controler

import (
	"fmt"
	"net/http"
)

func Server() {
	fmt.Println("GO to: http://127.0.0.1:8080/auth")
	mux := http.NewServeMux()

	mux.HandleFunc("/auth", Auth)
	mux.HandleFunc("/apply", Home)

	fs := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", mux)
}
