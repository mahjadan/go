package main

import (
	"github.com/gorilla/mux"
	"go-proto-poc/handler"
	"log"
	"net/http"
)

var port = ":8080"

func main() {
	handler := handler.Handler{}
	router := mux.NewRouter()
	router.HandleFunc("/health", HandleHealth).Methods(http.MethodGet)
	router.HandleFunc("/authentication_clients", handler.CreateAuthenticationClientJSON).Methods(http.MethodPost)

	log.Println("[GO-JSON] server starting on port: ", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func HandleHealth(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("{\"status\": \"OK\"}"))
}
