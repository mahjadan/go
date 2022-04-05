# proto-POC
gandauth POC to understand protobuf 

## Run:
`go run go-proto/main.go`

`go run go-json/main.go`

Or

`docker-compose up`

Then

`cd ruby-proto && ruby poc_client.rb`

## What was not implemented:
- Logging 
- Instrumentation

### Current design:
In A3s-api the generated proto request is used inside the service which use another library called `platformmodel`
which is also a proto message with some functionality to transform the proto message into an interface back and forth 
and also validate this interface based on predefined requestDefinitions and validationRules.

The same object (interface) passed from the service to the repository to be persisted and in some cases it is altered (to add more fields such as `created_at`)

This design is tightly coupled to the protobuf protocol and to the `platformmodel` library.
And it's not clear to who is going to read the code what are the business logic and what data need to be persisted because everything is mixed.

### Why this re-design:
The idea is to decouple our transport layer from the service and repository layers, and to implement this
we needed to restrict the use of proto structs from being passed down to the service and the repository.

This will make the business logic more clear and what data is being persisted is well-defined.
This will also allow us to use different repository or different transport layer without touching the service layer code.

Steps that were followed:
- replicate the same request message of a3s-api POST /authentication_clients to create a client.
- generate the proto files (golang,ruby)
- construct a handler that receives this protobuf request.
- construct a service model that will be used inside the service and map the request to this model.
- construct a service and repository
- configure the validator (playground-validator) to work with json-names and pass it to the service.
- validate the model inside the service.
- construct error model to transport specific error from the repository and service layer back to the transport layer (handler)
- replicate the same response message of a3s-api.
- construct a error proto message.
- generate the response proto files (golang,ruby)



## Links :
- [Protobuf](https://developers.google.com/protocol-buffers/docs/overview)
- [Grpc golang](https://grpc.io/docs/languages/go/quickstart/)
- [hexagonal architecture clean code](https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/)
