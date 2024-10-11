package main

import (
	"log"
	"net/http"

	"github.com/ACordoba15/be-user-maintenance/db"
	_ "github.com/ACordoba15/be-user-maintenance/docs"
	"github.com/ACordoba15/be-user-maintenance/internal/domain/repository"
	"github.com/ACordoba15/be-user-maintenance/internal/middleware"
	"github.com/ACordoba15/be-user-maintenance/internal/routes"
	"github.com/ACordoba15/be-user-maintenance/internal/usecase"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API de ejemplo con Swagger y Gorilla Mux
// @version 1.0
// @description Esta es una API en Go documentada con Swagger y usando Gorilla Mux.

// @host localhost:8000
// @BasePath /api/
func main() {
	// Conectar a la base de datos
	database, err := db.DBConnection()

	if err != nil {
		log.Fatal("Error inicializando la base de datos: ", err)
	}

	// Migrar los modelos
	err = db.MigrateDB(database)
	if err != nil {
		log.Fatal("Error al migrar modelos: ", err)
	}

	// Crear el repositorio e inyectarlo en el caso de uso
	userRepo := repository.NewUserRepository(database)
	userService := usecase.NewUserService(userRepo)

	recordRepo := repository.NewRecordRepository(database) // Asegúrate de tener este repositorio
	recordService := usecase.NewRecordService(recordRepo)

	// Crear router
	r := mux.NewRouter()

	// Endpoint de Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Registrar rutas
	routes.RegisterRoutes(r, userService, recordService)

	// Envolver con middleware (como CORS)
	handler := middleware.EnableCORS(r)

	// Iniciar servidor
	log.Fatal(http.ListenAndServe(":8000", handler))
}
