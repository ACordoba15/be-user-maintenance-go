package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ACordoba15/be-user-maintenance/internal/domain/models"
	"github.com/ACordoba15/be-user-maintenance/internal/usecase"
	"github.com/gorilla/mux"
)

// Dependencia de UserService (caso de uso)
var userService *usecase.UserService

func RegisterUserRoutes(r *mux.Router, us *usecase.UserService) {
	userService = us

	r.HandleFunc("/api/user/all", GetUsersHandler(userService)).Methods("GET")
	r.HandleFunc("/api/user/{id}", GetUserHandler(userService)).Methods("GET")
	r.HandleFunc("/api/user", PostUserHandler((userService))).Methods("POST")
	r.HandleFunc("/api/user/login", LoginUserHandler((userService))).Methods("POST")
	r.HandleFunc("/api/user", PutUserHandler((userService))).Methods("PUT")
	r.HandleFunc("/api/user/{id}", DeleteUserHandler((userService))).Methods("DELETE")
}

// GetUsersHandler obtiene todos los registros de la base de datos.
// @Summary Obtiene todos los registros
// @Description Retorna una lista de todos los registros almacenados en la base de datos.
// @Tags user
// @Produce  json
// @Success 200 {array} models.User
// @Router /user/all [get]
func GetUsersHandler(userService *usecase.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userService.GetAll()

		if err != nil {
			http.Error(w, "Error al obtener los usuarios", http.StatusInternalServerError)
			return
		}

		// Serializa los usuarios a JSON y envía la respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
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
func GetUserHandler(userService *usecase.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			http.Error(w, "Request inválido", http.StatusBadRequest)
			return
		}

		user, err := userService.GetById(id)

		if err != nil {
			http.Error(w, "Error al obtener el usuario", http.StatusInternalServerError)
			return
		}

		if user.ID == 0 {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}

		// Serializa los usuarios a JSON y envía la respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
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
func LoginUserHandler(userService *usecase.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login models.User
		err := json.NewDecoder(r.Body).Decode(&login)

		if err != nil {
			http.Error(w, "Request inválido", http.StatusBadRequest)
			return
		}

		user, err := userService.Login(login.Username, login.Password)

		if err != nil {
			http.Error(w, "Nombre de usuario o contraseña incorrecta", http.StatusNotFound)
			return
		}

		// Serializa los usuarios a JSON y envía la respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
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
func PostUserHandler(userService *usecase.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Request inválido", http.StatusBadRequest)
			return
		}

		newUser, err := userService.AddUser(user)

		if err != nil {
			http.Error(w, "Error al agregar el usuario.", http.StatusInternalServerError)
			return
		}

		// Serializa los usuarios a JSON y envía la respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newUser)
	}
}

// PutUserHandler actualiza un registro.
// @Summary Actualiza un registro
// @Description Actualiza la información de un registro.
// @Tags user
// @Produce plain
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad Request"
// @Router /user [put]
func PutUserHandler(userService *usecase.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Request inválido", http.StatusBadRequest)
			return
		}

		fmt.Println(user.Username, user.Password)
		updatedUser, err := userService.UpdateUser(user.Username, user.Password)

		if updatedUser.ID == 0 {
			http.Error(w, "Usuario no encontrado!", http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, "Error al actualizar el usuario.", http.StatusInternalServerError)
			return
		}

		// Serializa los usuarios a JSON y envía la respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedUser)
	}
}

// DeleteUserHandler elimina un registro por su ID.
// @Summary Elimina un registro por ID
// @Description Realiza el borrado lógico de un registro específico.
// @Tags user
// @Param id path int true "ID del registro"
// @Success 204 No Content
// @Failure 404 {string} string "User Not Found"
// @Router /user/{id} [delete]
func DeleteUserHandler(userService *usecase.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			http.Error(w, "Request inválido", http.StatusBadRequest)
			return
		}

		err = userService.DeleteUser(id)
		if err != nil {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}

		// Serializa los usuarios a JSON y envía la respuesta
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}
