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

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.ListenAndServe(":8080", mux)
}
