package repository

import (
	"github.com/google/uuid"
	"go-proto-poc/pkg/model"
	"time"
)

// convert the domain model to persistence model and return the clientId
func toPersistence(authClient model.AuthenticationClient) (interface{}, string) {

	switch authClient.AuthenticationClientType {
	case model.JWT:
		return JWTAuthenticationClient{
			AuthenticationClientType: authClient.AuthenticationClientType,
			ClientId:                 authClient.ClientId,
			PlatformAccountId:        authClient.PlatformAccountId,
			ClientSecret:             authClient.ClientSecret,
			ApplicationUuid:          authClient.ApplicationUuid,
			Enabled:                  authClient.Enabled,
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
		}, authClient.ClientId

	case model.API_KEY:
		return ApiKeyAuthenticationClient{
			AuthenticationClientType: authClient.AuthenticationClientType,
			ClientId:                 authClient.ClientId,
			PlatformAccountId:        authClient.PlatformAccountId,
			ClientSecret:             authClient.ClientSecret,
			ApplicationUuid:          authClient.ApplicationUuid,
			Enabled:                  authClient.Enabled,
			UserId:                   authClient.UserId,
			ApiKey:                   authClient.ApiKey,
			ApiKeyName:               authClient.ApiKeyName,
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
		}, authClient.ClientId

	case model.PRIVATE_TOKEN, model.PUBLIC_TOKEN:
		client := TokenAuthenticationClient{
			AuthenticationClientType: authClient.AuthenticationClientType,
			PlatformAccountId:        authClient.PlatformAccountId,
			ApplicationUuid:          authClient.ApplicationUuid,
			Enabled:                  authClient.Enabled,
			ApiKey:                   authClient.ApiKey,
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
		}
		client.ClientId = uuid.New().String()
		if client.ApiKey == "" {
			client.ApiKey = uuid.New().String()
		}
		return client, client.ClientId

	default:
		return nil, ""
	}
}

func toModel(persistence interface{}) model.AuthenticationClient {

	switch authClient := persistence.(type) {
	case ApiKeyAuthenticationClient:
		return model.AuthenticationClient{
			AuthenticationClientType: authClient.AuthenticationClientType,
			ClientId:                 authClient.ClientId,
			PlatformAccountId:        authClient.PlatformAccountId,
			ClientSecret:             authClient.ClientSecret,
			ApplicationUuid:          authClient.ApplicationUuid,
			Enabled:                  authClient.Enabled,
			UserId:                   authClient.UserId,
			ApiKey:                   authClient.ApiKey,
			ApiKeyName:               authClient.ApiKeyName,
			CreatedAt:                authClient.CreatedAt,
			UpdatedAt:                authClient.UpdatedAt,
		}

	case TokenAuthenticationClient:
		return model.AuthenticationClient{
			AuthenticationClientType: authClient.AuthenticationClientType,
			ClientId:                 authClient.ClientId,
			PlatformAccountId:        authClient.PlatformAccountId,
			ApplicationUuid:          authClient.ApplicationUuid,
			Enabled:                  authClient.Enabled,
			ApiKey:                   authClient.ApiKey,
			CreatedAt:                authClient.CreatedAt,
			UpdatedAt:                authClient.UpdatedAt,
		}

	case JWTAuthenticationClient:
		return model.AuthenticationClient{
			AuthenticationClientType: authClient.AuthenticationClientType,
			ClientId:                 authClient.ClientId,
			PlatformAccountId:        authClient.PlatformAccountId,
			ClientSecret:             authClient.ClientSecret,
			ApplicationUuid:          authClient.ApplicationUuid,
			Enabled:                  authClient.Enabled,
			CreatedAt:                authClient.CreatedAt,
			UpdatedAt:                authClient.UpdatedAt,
		}

	default:
		return model.AuthenticationClient{}
	}
}
