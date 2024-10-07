package main

import (
	"net/http"

	"github.com/ACordoba15/be-user-maintenance/db"
	_ "github.com/ACordoba15/be-user-maintenance/docs"
	"github.com/ACordoba15/be-user-maintenance/models"
	"github.com/ACordoba15/be-user-maintenance/routes"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API de ejemplo con Swagger y Gorilla Mux
// @version 1.0
// @description Esta es una API en Go documentada con Swagger y usando Gorilla Mux.

// @host localhost:8000
// @BasePath /api/
func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.Record{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	// Endpoint de Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/", routes.HomeHandler)

	// USER
	r.HandleFunc("/api/user/all", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/api/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/api/user", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/api/user/login", routes.LoginUserHandler).Methods("POST")
	r.HandleFunc("/api/user/{id}", routes.PutUserHandler).Methods("PUT")
	r.HandleFunc("/api/user/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// RECORD
	r.HandleFunc("/api/record/all", routes.GetRecordHandler).Methods("GET")
	r.HandleFunc("/api/record/{id}", routes.GetRecordHandler).Methods("GET")
	r.HandleFunc("/api/record", routes.PostRecordHandler).Methods("POST")
	r.HandleFunc("/api/record/{id}", routes.PutRecordHandler).Methods("PUT")
	r.HandleFunc("/api/record/{id}", routes.DeleteRecordHandler).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
