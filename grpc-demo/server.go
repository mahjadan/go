package main

import (
	"context"
	"flag"
	"fmt"
	mygrpc "github.com/mahjadan/go/grpc-demo/pkg/grpc"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type server struct {
	mygrpc.UnimplementedCpfValidatorServer
}

func main() {
	port := flag.Int("port", 8080, "grpc server port")
	flag.Parse()

	srv := NewServer()

	grpcServer := grpc.NewServer()
	mygrpc.RegisterCpfValidatorServer(grpcServer, srv)

	address := ":" + strconv.Itoa(*port)
	listen, err2 := net.Listen("tcp", address)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("listening on :", address)
	grpcServer.Serve(listen)

}

func NewServer() *server {
	return &server{}
}

func (s *server) Validate(ctx context.Context, req *mygrpc.CpfRequest) (*mygrpc.CpfResponse, error) {
	fmt.Println("[SERVER] receive request...")
	var constraints []*mygrpc.ConstraintResponse
	c := mygrpc.ConstraintResponse{
		Name:   req.Validations[0].Name,
		Value:  req.Validations[0].Value,
		Result: true,
	}
	constraints = append(constraints, &c)
	return &mygrpc.CpfResponse{
		Validations: constraints,
	}, nil
}
