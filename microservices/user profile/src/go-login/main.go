package main

import (
	"go-login/controller"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", controller.PingHandler).
	Methods("GET")
	r.HandleFunc("/register", controller.RegisterHandler).
		Methods("POST","OPTIONS")
	r.HandleFunc("/login", controller.LoginHandler).
		Methods("POST","OPTIONS")
	r.HandleFunc("/profile", controller.ProfileHandler).
		Methods("GET")
	
	r.Use(mux.CORSMethodMiddleware(r))
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(r)))
}