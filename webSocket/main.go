package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"webSocket/pubsub"
)

var ps pubsub.PubSub
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


func main() {
	fmt.Println("hello world.")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ws",webSocket)
log.Fatal(http.ListenAndServe(":8080",nil))
}


func webSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client:= pubsub.Client{
		Id: generateID(),
		Conn: conn,
	}
	fmt.Println("new Client connected: ",client.Id)
	ps.AddClient(&client)

	for {
		fmt.Println("total subscriptions: ",len(ps.Subscriptions))
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("something went wrong while reading: %s :",err)
			ps.RemoveClient(client.Id)
			return
		}
		ps.HandleMessage(client,messageType,p)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"static")
}

func generateID() string{
	return uuid.NewV4().String()
}