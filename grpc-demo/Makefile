install-grpc:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/prtoc-gen-go

gen:
	protoc --go_out=. --go_opt=module=github.com/mahjadan/go/grpc-demo \
 	--go-grpc_out=. --go-grpc_opt=module=github.com/mahjadan/go/grpc-demo \
 	 --grpc-gateway_out=pkg proto/*.proto

clean:
	rm -rf pkg/*

server:
	go run server.go -port 8080

client:
	go run client.go -port 8080

task_server:
	go run task/server.go -port 8080

task_client:
	go run task/client.go -port 8080
