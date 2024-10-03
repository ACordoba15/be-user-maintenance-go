package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ACordoba15/be-user-maintenance/db"
	"github.com/ACordoba15/be-user-maintenance/models"
	"github.com/gorilla/mux"
)

func GetRecordsHandler(w http.ResponseWriter, r *http.Request) {
	var records []models.Record
	db.DB.Find(&records)
	json.NewEncoder(w).Encode(&records)
}

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

func PostRecordHandler(w http.ResponseWriter, r *http.Request) {
	var record models.Record
	json.NewDecoder(r.Body).Decode(&record)

	newrecord := db.DB.Create(&record)
	err := newrecord.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&record)
}

func PutRecordHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Put Record"))
}

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
	// db.DB.Unscoped().Delete(&record) // Borrado físico
	w.WriteHeader(http.StatusOK)
}
