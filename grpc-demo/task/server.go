package main

import (
	"context"
	"flag"
	"fmt"
	mygrpc "github.com/mahjadan/go/grpc-demo/pkg/grpc"
	repo2 "github.com/mahjadan/go/grpc-demo/task/repo"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type listServer struct {
	mygrpc.UnimplementedTasksServer
	store repo2.Store
}

func main() {
	port := flag.Int("port", 8080, "grpc server port")
	flag.Parse()

	srv := NewListServer()

	grpcServer := grpc.NewServer()
	mygrpc.RegisterTasksServer(grpcServer, srv)

	address := ":" + strconv.Itoa(*port)
	listen, err2 := net.Listen("tcp", address)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("listening on :", address)
	log.Fatal(grpcServer.Serve(listen))

}

func NewListServer() *listServer {
	return &listServer{
		store: repo2.NewInMemoryStore(),
	}
}

func (s *listServer) List(ctx context.Context, v *mygrpc.Void) (*mygrpc.TaskList, error) {
	fmt.Println("[SERVER] receive list request...")
	taskList := s.store.GetAll()
	return &taskList, nil
}

func (s *listServer) Add(ctx context.Context, newTask *mygrpc.NewTask) (*mygrpc.Task, error) {
	fmt.Println("[SERVER] receive add request...")
	task := mygrpc.Task{
		Name: newTask.Name,
		Done: false,
	}
	s.store.Save(task)
	return &task, nil
}
