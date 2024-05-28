package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Imjowend/citas-golang/server"
)

type CitaResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func InsertCita(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CitaResponse{
			Message: "Welcome to Cita",
			Status:  true,
		})
	}
}

func UpdateCitaByDNI(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CitaResponse{
			Message: "Welcome to Cita",
			Status:  true,
		})
	}
}
func GetCitas(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CitaResponse{
			Message: "Welcome to Cita",
			Status:  true,
		})
	}
}
func GetCitaByDNI(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CitaResponse{
			Message: "Welcome to Cita",
			Status:  true,
		})
	}
}
func DeleteCitaByDNI(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CitaResponse{
			Message: "Welcome to Cita",
			Status:  true,
		})
	}
}
