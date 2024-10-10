package routes

import (
	"github.com/ACordoba15/be-user-maintenance/internal/usecase"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, userService *usecase.UserService, recordService *usecase.RecordService) {
	// Rutas Home
	r.HandleFunc("/", HomeHandler).Methods("GET")

	// Rutas de usuario
	RegisterUserRoutes(r, userService)

	// Rutas de registros
	RegisterRecordRoutes(r, recordService)
}
