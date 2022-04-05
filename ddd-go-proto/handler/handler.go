package handler

import (
	"go-proto-poc/pkg/model"
	"go-proto-poc/pkg/service"
	pb "go-proto-poc/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
	Service service.Service
}

func NewHandler(service service.Service) Handler {
	return Handler{
		Service: service,
	}
}
func (h Handler) CreateAuthenticationClientPROTO(w http.ResponseWriter, r *http.Request) {
	request := pb.CreateAuthenticationClient{}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		h.writeError(w, codes.InvalidArgument, err)
		return
	}
	err = proto.Unmarshal(bodyBytes, &request)
	if err != nil {
		log.Println(err)
		h.writeError(w, codes.InvalidArgument, err)
		return
	}

	modelResponse, err := h.Service.CreateAuthenticationClient(r.Context(), requestToModel(request))
	if err != nil {
		if e, ok := err.(*model.Error); ok {
			if e.ErrorCode == "Validations" {
				h.writeError(w, codes.InvalidArgument, err)
				return
			}
			if e.ErrorCode == "INTERNAL" { // example of errorCode that we can return in our service to help return a specific statusCode
				// refactor to pass the code or the status code ?
				h.writeError(w, codes.Internal, err)
				return
			}
		} else {
			log.Println(err)
			h.writeError(w, codes.Internal, err)
			return
		}
	}
	response := toProtoResponse(modelResponse)
	responseBytes, err := proto.Marshal(&response)
	if err != nil {
		log.Println(err)
		h.writeError(w, codes.InvalidArgument, err)
		return
	}
	h.writeResponse(w, http.StatusOK, responseBytes)
}

func (h Handler) writeError(w http.ResponseWriter, code codes.Code, err error) {
	errResponse := toErrorProto(code, err)
	responseBytes, err := proto.Marshal(&errResponse)
	if err != nil {
		log.Println(err)
	}
	h.writeResponse(w, codeToHttpStatus(code), responseBytes)
}

func (h Handler) writeResponse(w http.ResponseWriter, status int, responseBytes []byte) {
	w.Header().Set("Content-Type", "application/protobuf")
	w.WriteHeader(status)
	w.Write(responseBytes)
}
