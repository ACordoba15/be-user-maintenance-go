package main

import (
	"net/http"

	"github.com/ACordoba15/be-user-maintenance/db"
	_ "github.com/ACordoba15/be-user-maintenance/docs"
	"github.com/ACordoba15/be-user-maintenance/models"
	"github.com/ACordoba15/be-user-maintenance/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                            // Permitir todos los orígenes, puedes especificar orígenes específicos aquí.
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"}, // Métodos permitidos
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// USER
	r.HandleFunc("/api/user/all", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/api/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/api/user", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/api/user/login", routes.LoginUserHandler).Methods("POST")
	r.HandleFunc("/api/user", routes.PutUserHandler).Methods("PUT")
	r.HandleFunc("/api/user/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// RECORD
	r.HandleFunc("/api/record/all", routes.GetRecordsHandler).Methods("GET")
	r.HandleFunc("/api/record/{id}", routes.GetRecordHandler).Methods("GET")
	r.HandleFunc("/api/record", routes.PostRecordHandler).Methods("POST")
	r.HandleFunc("/api/record/{id}", routes.PutRecordHandler).Methods("PUT")
	r.HandleFunc("/api/record/{id}", routes.DeleteRecordHandler).Methods("DELETE")

	// Envolver el router con el middleware de CORS
	handler := c.Handler(r)

	http.ListenAndServe(":8000", handler)
}
