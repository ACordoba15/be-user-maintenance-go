package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ACordoba15/be-user-maintenance/db"
	"github.com/ACordoba15/be-user-maintenance/models"
	"github.com/gorilla/mux"
)

// GetRecordsHandler obtiene todos los registros de la base de datos.
// @Summary Obtiene todos los registros
// @Description Retorna una lista de todos los registros almacenados en la base de datos.
// @Tags record
// @Produce  json
// @Success 200 {array} models.Record
// @Router /record/all [get]
func GetRecordsHandler(w http.ResponseWriter, r *http.Request) {
	var records []models.Record
	db.DB.Find(&records)
	json.NewEncoder(w).Encode(&records)
}

// GetRecordHandler obtiene un registro por su ID.
// @Summary Obtiene un registro por ID
// @Description Retorna un registro específico basado en el ID proporcionado.
// @Tags record
// @Param id path int true "ID del registro"
// @Produce  json
// @Success 200 {object} models.Record
// @Failure 404 {string} string "Record Not Found"
// @Router /record/{id} [get]
func GetRecordHandler(w http.ResponseWriter, r *http.Request) {
	var record models.Record
	params := mux.Vars(r)
	db.DB.First(&record, params["id"])

	if record.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Record Not Found"))
		return
	}

	json.NewEncoder(w).Encode(&record)
}

// PostRecordHandler crea un nuevo registro.
// @Summary Crea un nuevo registro
// @Description Agrega un nuevo registro a la base de datos.
// @Tags record
// @Accept  json
// @Produce  json
// @Param record body models.Record true "Información del nuevo registro"
// @Success 200 {object} models.Record
// @Failure 400 {string} string "Bad Request"
// @Router /record [post]
func PostRecordHandler(w http.ResponseWriter, r *http.Request) {
	var record models.Record
	json.NewDecoder(r.Body).Decode(&record)
	defer r.Body.Close()

	newrecord := db.DB.Create(&record)
	err := newrecord.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&record)
}

// PutRecordHandler actualiza un registro.
// @Summary Actualiza un registro
// @Description Actualiza la información de un registro.
// @Tags record
// @Produce plain
// @Success 200 {object} models.Record
// @Failure 400 {string} string "Bad Request"
// @Router /record/{id} [put]
func PutRecordHandler(w http.ResponseWriter, r *http.Request) {
	var record models.Record
	var newRecord models.Record
	params := mux.Vars(r) // Obtiene los params
	db.DB.First(&record, params["id"])

	defer r.Body.Close() // Libera recursos

	if record.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Record Not Found"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&newRecord)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}
	record.Action = newRecord.Action
	db.DB.Save(&record)
	json.NewEncoder(w).Encode(&record)
}

// DeleteRecordHandler elimina un registro por su ID.
// @Summary Elimina un registro por ID
// @Description Realiza el borrado lógico de un registro específico.
// @Tags record
// @Param id path int true "ID del registro"
// @Success 200 {string} string "Record Deleted"
// @Failure 404 {string} string "Record Not Found"
// @Router /record/{id} [delete]
func DeleteRecordHandler(w http.ResponseWriter, r *http.Request) {
	var record models.Record
	params := mux.Vars(r)
	db.DB.First(&record, params["id"])

	if record.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Record Not Found"))
		return
	}

	db.DB.Delete(&record) // Borrado lógico
	w.WriteHeader(http.StatusOK)
}
