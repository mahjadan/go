generate:
	protoc --proto_path=./protos \
            --go-grpc_out=./protos \
            --go-grpc_opt=paths=source_relative \
            --go_out=paths=source_relative:./protos \
            --ruby_out=paths=source_relative:./protos/ \
                protos/*.proto
