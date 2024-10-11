package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ACordoba15/be-user-maintenance/internal/domain/models"
	"github.com/ACordoba15/be-user-maintenance/internal/usecase"
	"github.com/gorilla/mux"
)

var recordService *usecase.RecordService

func RegisterRecordRoutes(r *mux.Router, rs *usecase.RecordService) {
	recordService = rs

	r.HandleFunc("/api/record/all", GetRecordsHandler(recordService)).Methods("GET")
	r.HandleFunc("/api/record/{id}", GetRecordHandler(recordService)).Methods("GET")
	r.HandleFunc("/api/record", PostRecordHandler(recordService)).Methods("POST")
	r.HandleFunc("/api/record/{id}", PutRecordHandler(recordService)).Methods("PUT")
	r.HandleFunc("/api/record/{id}", DeleteRecordHandler(recordService)).Methods("DELETE")
}

// GetRecordsHandler obtiene todos los registros de la base de datos.
// @Summary Obtiene todos los registros
// @Description Retorna una lista de todos los registros almacenados en la base de datos.
// @Tags record
// @Produce  json
// @Success 200 {array} models.Record
// @Router /record/all [get]
func GetRecordsHandler(recordService *usecase.RecordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		records, err := recordService.GetAll()

		if err != nil {
			http.Error(w, "Error al obtener los registros", http.StatusInternalServerError)
			return
		}

		// Serializa los usuarios a JSON y envía la respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(records)
	}
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
func GetRecordHandler(recordService *usecase.RecordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			http.Error(w, "Request inválido", http.StatusBadRequest)
			return
		}

		record, err := recordService.GetById(id)

		if err != nil {
			http.Error(w, "Error al obtener el registro", http.StatusInternalServerError)
			return
		}

		if record.ID == 0 {
			http.Error(w, "Registro no encontrado", http.StatusNotFound)
			return
		}

		// Serializa los usuarios a JSON y envía la respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(record)
	}
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
func PostRecordHandler(recordService *usecase.RecordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record models.Record
		err := json.NewDecoder(r.Body).Decode(&record)

		if err != nil {
			http.Error(w, "Request inválido", http.StatusBadRequest)
			return
		}

		newRecord, err := recordService.AddRecord(record)

		if err != nil {
			http.Error(w, "Error al agregar el registro.", http.StatusInternalServerError)
			return
		}

		// Serializa los usuarios a JSON y envía la respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newRecord)
	}
}

// PutRecordHandler actualiza un registro.
// @Summary Actualiza un registro
// @Description Actualiza la información de un registro.
// @Tags record
// @Produce plain
// @Success 200 {object} models.Record
// @Failure 400 {string} string "Bad Request"
// @Router /record/{id} [put]
func PutRecordHandler(recordService *usecase.RecordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record models.Record
		err := json.NewDecoder(r.Body).Decode(&record)

		if err != nil {
			http.Error(w, "Request inválido", http.StatusBadRequest)
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			http.Error(w, "Request inválido", http.StatusBadRequest)
			return
		}

		updatedRecord, err := recordService.UpdateRecord(record, id)

		if err != nil {
			http.Error(w, "Error al actualizar el registro.", http.StatusInternalServerError)
			return
		}

		// Serializa los usuarios a JSON y envía la respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedRecord)
	}
}

// DeleteRecordHandler elimina un registro por su ID.
// @Summary Elimina un registro por ID
// @Description Realiza el borrado lógico de un registro específico.
// @Tags record
// @Param id path int true "ID del registro"
// @Success 204 No Content
// @Failure 404 {string} string "Record Not Found"
// @Router /record/{id} [delete]
func DeleteRecordHandler(recordService *usecase.RecordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			http.Error(w, "Request inválido", http.StatusBadRequest)
			return
		}

		err = recordService.DeleteRecord(id)

		if err != nil {
			http.Error(w, "Registro no encontrado", http.StatusNotFound)
			return
		}

		// Serializa los usuarios a JSON y envía la respuesta
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}
