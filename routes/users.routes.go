package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ACordoba15/be-user-maintenance/db"
	"github.com/ACordoba15/be-user-maintenance/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
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
// @Summary Login de usuario
// @Description Valida un usuari registrado.
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.User true "Información del nuevo registro"
// @Success 200 {object} models.Record
// @Failure 400 {string} string "Bad Request"
// @Router /user/login [post]
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var userLogin models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte("Invalid request payload"))
		return
	}

	fmt.Printf("user: %s, pass: %s", user.Username, user.Password)
	// Buscar el usuario por nombre de usuario y contraseña
	err = db.DB.Where(&models.User{Username: user.Username, Password: user.Password}).First(&userLogin).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound) // 404
			w.Write([]byte("User Not Found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError) // 500
			w.Write([]byte("Internal Server Error"))
		}
		return
	}

	if userLogin.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("User Not Found"))
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(&userLogin)
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
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad Request"
// @Router /user [put]
func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var userUpdated models.User
	err := json.NewDecoder(r.Body).Decode(&userUpdated)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte("Invalid request payload"))
		return
	}

	// Buscar el usuario por nombre de usuario y contraseña
	err = db.DB.Where(&models.User{Username: user.Username, Password: user.Password}).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound) // 404
			w.Write([]byte("User Not Found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError) // 500
			w.Write([]byte("Internal Server Error"))
		}
		return
	}

	user.Password = userUpdated.Password
	db.DB.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

// DeleteUserHandler elimina un registro por su ID.
// @Summary Elimina un registro por ID
// @Description Realiza el borrado lógico de un registro específico.
// @Tags user
// @Param id path int true "ID del registro"
// @Success 204 No Content
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
	w.WriteHeader(http.StatusNoContent)
}
