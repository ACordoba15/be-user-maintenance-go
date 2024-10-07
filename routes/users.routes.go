package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ACordoba15/be-user-maintenance/db"
	"github.com/ACordoba15/be-user-maintenance/models"
	"github.com/gorilla/mux"
)

// GetUsersHandler obtiene todos los registros de la base de datos.
// @Summary Obtiene todos los registros
// @Description Retorna una lista de todos los registros almacenados en la base de datos.
// @Tags user
// @Produce  json
// @Success 200 {array} models.User
// @Router /user/all [get]
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

// GetUserHandler obtiene un registro por su ID.
// @Summary Obtiene un registro por ID
// @Description Retorna un registro específico basado en el ID proporcionado.
// @Tags user
// @Param id path int true "ID del registro"
// @Produce  json
// @Success 200 {object} models.User
// @Failure 404 {string} string "Record Not Found"
// @Router /user/{id} [get]
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

// LoginUserHandler crea un nuevo registro.
// @Summary Crea un nuevo registro
// @Description Agrega un nuevo registro a la base de datos.
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.User true "Información del nuevo registro"
// @Success 200 {object} models.Record
// @Failure 400 {string} string "Bad Request"
// @Router /user/login [post]
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

// PostUserHandler crea un nuevo registro.
// @Summary Crea un nuevo registro
// @Description Agrega un nuevo registro a la base de datos.
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.User true "Información del nuevo registro"
// @Success 200 {object} models.Record
// @Failure 400 {string} string "Bad Request"
// @Router /user [post]
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	newUser := db.DB.Create(&user)
	err := newUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

// PutUserHandler actualiza un registro.
// @Summary Actualiza un registro
// @Description Actualiza la información de un registro.
// @Tags user
// @Produce plain
// @Router /user/{id} [put]
func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Put User"))
}

// DeleteUserHandler elimina un registro por su ID.
// @Summary Elimina un registro por ID
// @Description Realiza el borrado lógico de un registro específico.
// @Tags user
// @Param id path int true "ID del registro"
// @Success 200 {string} string "User Deleted"
// @Failure 404 {string} string "User Not Found"
// @Router /user/{id} [delete]
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("User Not Found"))
		return
	}

	db.DB.Delete(&user) // Borrado lógico
	// db.DB.Unscoped().Delete(&user) // Borrado físico
	w.WriteHeader(http.StatusOK)
}
