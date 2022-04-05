package repository

import (
	"go-proto-poc/pkg/model"
	"time"
)

type TokenAuthenticationClient struct {
	AuthenticationClientType model.AuthenticationClientType `firestore:"authentication_client_type"`
	ClientId                 string                         `firestore:"client_id"`
	PlatformAccountId        int64                          `firestore:"platform_account_id"`
	ApplicationUuid          string                         `firestore:"application_uuid"`
	Enabled                  bool                           `firestore:"enabled"`
	ApiKey                   string                         `firestore:"api_key"`
	CreatedAt                time.Time                      `firestore:"created_at"`
	UpdatedAt                time.Time                      `firestore:"updated_at"`
}
type JWTAuthenticationClient struct {
	AuthenticationClientType model.AuthenticationClientType `firestore:"authentication_client_type"`
	ClientId                 string                         `firestore:"client_id"`
	PlatformAccountId        int64                          `firestore:"platform_account_id"`
	ClientSecret             string                         `firestore:"client_secret"`
	ApplicationUuid          string                         `firestore:"application_uuid"`
	Enabled                  bool                           `firestore:"enabled"`
	CreatedAt                time.Time                      `firestore:"created_at"`
	UpdatedAt                time.Time                      `firestore:"updated_at"`
}
type ApiKeyAuthenticationClient struct {
	AuthenticationClientType model.AuthenticationClientType `firestore:"authentication_client_type"`
	ClientId                 string                         `firestore:"client_id"`
	PlatformAccountId        int64                          `firestore:"platform_account_id"`
	ClientSecret             string                         `firestore:"client_secret"`
	ApplicationUuid          string                         `firestore:"application_uuid"`
	Enabled                  bool                           `firestore:"enabled"`
	UserId                   int64                          `firestore:"user_id"`
	ApiKey                   string                         `firestore:"api_key"`
	ApiKeyName               string                         `firestore:"api_key_name"`
	CreatedAt                time.Time                      `firestore:"created_at"`
	UpdatedAt                time.Time                      `firestore:"updated_at"`
}
