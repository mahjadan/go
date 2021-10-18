package main

import (
	"context"
	"flag"
	"fmt"
	mygrpc "github.com/mahjadan/go/grpc-demo/pkg/grpc"
	"google.golang.org/grpc"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	port := flag.Int("port", 8080, "port of the grpc server to connect to")
	flag.Parse()
	fmt.Println("connecting to :", *port)
	conn, err := grpc.Dial(":"+strconv.Itoa(*port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := mygrpc.NewTasksClient(conn)

	for {
		var constraints []*mygrpc.Constraint
		c := mygrpc.Constraint{
			Name:  "HAS_EMAIL",
			Value: "test@gmail.com",
		}
		constraints = append(constraints, &c)
		list, err := client.List(context.Background(), &mygrpc.Void{})
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
		fmt.Println("LIST_RESPONSE: ", list)
		time.Sleep(time.Second)
		go func() {
			ticker := time.NewTicker(2 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					client.Add(context.Background(), &mygrpc.NewTask{Name: "task-" + strconv.Itoa(rand.Intn(10))})
				}
			}
		}()
	}
}
