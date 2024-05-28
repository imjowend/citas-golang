package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Imjowend/citas-golang/handler"
	"github.com/Imjowend/citas-golang/middleware"
	"github.com/Imjowend/citas-golang/server"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))
	r.HandleFunc("/signup", handler.SignUp(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handler.Login(s)).Methods(http.MethodPost)
	//Registrar Cita
	r.HandleFunc("/citas", handler.InsertCita(s)).Methods(http.MethodPost)
	//Modificar Cita
	r.HandleFunc("/citas/{id}", handler.UpdateCitaByDNI(s)).Methods(http.MethodPut)
	//Listar Cita (Dni o Todos)
	r.HandleFunc("/citas", handler.GetCitas(s)).Methods(http.MethodGet)
	r.HandleFunc("/citas/{id}", handler.GetCitaByDNI(s)).Methods(http.MethodGet)
	//Eliminar Cita (Dni)
	r.HandleFunc("/citas/{id}", handler.DeleteCitaByDNI(s)).Methods(http.MethodDelete)
}
