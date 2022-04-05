package handler

import (
	"go-proto-poc/pkg/model"
	"go-proto-poc/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
)

func requestToModel(request pb.CreateAuthenticationClient) model.AuthenticationClient {
	return model.AuthenticationClient{
		AuthenticationClientType: authTypeToModel(request.GetAuthenticationClientType()),
		ClientId:                 request.GetClientId(),
		PlatformAccountId:        request.GetPlatformAccountId(),
		ClientSecret:             request.GetClientSecret(),
		ApplicationUuid:          request.GetApplicationUuid(),
		Enabled:                  request.GetEnabled(),
		UserId:                   request.GetUserId(),
		ApiKey:                   request.GetApiKey(),
		ApiKeyName:               request.GetApiKeyName(),
	}
}

func toProtoResponse(response model.AuthenticationClient) pb.CreateAuthenticationClientResponse {
	client := pb.AuthenticationClient{
		AuthenticationClientType: authTypeToProto(response.AuthenticationClientType),
		ClientId:                 response.ClientId,
		PlatformAccountId:        response.PlatformAccountId,
		ClientSecret:             response.ClientSecret,
		ApplicationUuid:          response.ApplicationUuid,
		Enabled:                  response.Enabled,
		UserId:                   response.UserId,
		ApiKey:                   response.ApiKey,
		ApiKeyName:               response.ApiKeyName,
		CreatedAt:                timestamppb.New(response.CreatedAt),
		UpdatedAt:                timestamppb.New(response.UpdatedAt),
	}
	return pb.CreateAuthenticationClientResponse{
		AuthClient: &client,
	}
}

func authTypeToProto(authType model.AuthenticationClientType) pb.AuthenticationClientType {
	switch authType {
	case model.UNKNOWN:
		return pb.AuthenticationClientType_AUTHENTICATION_CLIENT_TYPE_UNSPECIFIED
	case model.JWT:
		return pb.AuthenticationClientType_JWT
	case model.PUBLIC_TOKEN:
		return pb.AuthenticationClientType_PUBLIC_TOKEN
	case model.API_KEY:
		return pb.AuthenticationClientType_API_KEY
	case model.PRIVATE_TOKEN:
		return pb.AuthenticationClientType_PRIVATE_TOKEN
	}
	return pb.AuthenticationClientType_AUTHENTICATION_CLIENT_TYPE_UNSPECIFIED
}
func authTypeToModel(authType pb.AuthenticationClientType) model.AuthenticationClientType {
	switch authType {
	case pb.AuthenticationClientType_AUTHENTICATION_CLIENT_TYPE_UNSPECIFIED:
		return model.UNKNOWN
	case pb.AuthenticationClientType_JWT:
		return model.JWT
	case pb.AuthenticationClientType_PUBLIC_TOKEN:
		return model.PUBLIC_TOKEN
	case pb.AuthenticationClientType_API_KEY:
		return model.API_KEY
	case pb.AuthenticationClientType_PRIVATE_TOKEN:
		return model.PRIVATE_TOKEN
	}
	return model.UNKNOWN
}
func codeToHttpStatus(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.InvalidArgument:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
