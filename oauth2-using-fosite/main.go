package main

import (
	"fmt"
	"log"
	"net/http"
	"oauth2-using-fosite/server"
)

func main() {

	server.RegisterHandlers()
	server.InitSessionStore()

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	port:= ":8080"
	fmt.Println("listening on port "+ port)
	log.Fatal(http.ListenAndServe( port, nil))

}
