package handler

import (
	"go-proto-poc/pkg/model"
	"go-proto-poc/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func toErrorProto(code codes.Code, err error) pb.CreateAuthenticationClientResponse {
	var st *status.Status
	if er, ok := err.(*model.Error); ok {
		ve := pb.RpcError{
			ErrorCode:        er.ErrorCode,
			Message:          er.Message,
			ValidationErrors: toRpcValidationError(er),
		}
		st, _ = status.New(code, er.Message).WithDetails(&ve)
	} else {
		st = status.New(code, err.Error())
	}
	return pb.CreateAuthenticationClientResponse{
		Status: st.Proto(),
	}
}
func toRpcValidationError(err *model.Error) []*pb.ValidationError {
	if err == nil {
		return nil
	}
	var validationErrors []*pb.ValidationError
	for i := 0; i < len(err.ValidationErrors); i++ {
		v := pb.ValidationError{
			Field:       err.ValidationErrors[i].Field,
			Restriction: string(err.ValidationErrors[i].Restriction),
			Message:     err.ValidationErrors[i].Message,
		}

		validationErrors = append(validationErrors, &v)
	}
	return validationErrors
}
