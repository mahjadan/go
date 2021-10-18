package main

import (
	"context"
	"flag"
	"fmt"
	mygrpc "github.com/mahjadan/go/grpc-demo/pkg/grpc"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

func main() {
	port := flag.Int("port", 8080, "port of the grpc server to connect to")
	flag.Parse()
	fmt.Println("connecting to :", *port)
	conn, err := grpc.Dial(":"+strconv.Itoa(*port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := mygrpc.NewCpfValidatorClient(conn)

	for {
		var constraints []*mygrpc.Constraint
		c := mygrpc.Constraint{
			Name:  "HAS_EMAIL",
			Value: "test@gmail.com",
		}
		constraints = append(constraints, &c)
		cpfResponse, err := client.Validate(context.Background(), &mygrpc.CpfRequest{

			Cpf:         "1111",
			Validations: constraints,
		})
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
		fmt.Println("RESPONSE: ", cpfResponse)
		time.Sleep(time.Second)
	}
}
