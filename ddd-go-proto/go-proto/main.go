package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go-proto-poc/handler"
	"go-proto-poc/pkg/repository"
	"go-proto-poc/pkg/service"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

var port = ":8888"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	firestoreClient, err := firestore.NewClient(ctx, os.Getenv("CLOUD_PROJECT_ID"))
	if err != nil {
		log.Fatal("firestore initialization error:", err)
	}

	repo := repository.NewRepository(firestoreClient)
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	service := service.NewService(validate, repo)
	h := handler.NewHandler(service)
	router := mux.NewRouter()
	router.HandleFunc("/authentication_clients", h.CreateAuthenticationClientPROTO).Methods(http.MethodPost)

	log.Println("[GO-PROTO] server starting on port: ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
