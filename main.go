package main

import (
	"net/http"

	"github.com/ACordoba15/be-user-maintenance/db"
	"github.com/ACordoba15/be-user-maintenance/models"
	"github.com/ACordoba15/be-user-maintenance/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.Record{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/api/user/all", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/api/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/api/user", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/api/user/login", routes.LoginUserHandler).Methods("POST")
	r.HandleFunc("/api/user/{id}", routes.PutUserHandler).Methods("PUT")
	r.HandleFunc("/api/user/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":5500", r)
}
