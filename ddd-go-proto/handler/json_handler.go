package handler

import (
	"encoding/json"
	"go-proto-poc/pkg/model"
	"log"
	"net/http"
)

func (h Handler) CreateAuthenticationClientJSON(w http.ResponseWriter, r *http.Request) {

	request := CreateAuthenticationRequestJSON{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(&request)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

type CreateAuthenticationRequestJSON struct {
	AuthenticationClientType model.AuthenticationClientType `json:"authentication_client_type"`
	ClientId                 string                         `json:"client_id,omitempty"`
	PlatformAccountId        int64                          `json:"platform_account_id,omitempty"`
	ClientSecret             string                         `json:"client_secret,omitempty"`
	ApplicationUuid          string                         `json:"application_uuid,omitempty"`
	Enabled                  bool                           `json:"enabled,omitempty"`
	UserId                   int64                          `json:"user_id,omitempty"`
	ApiKey                   string                         `json:"api_key,omitempty"`
	ApiKeyName               string                         `json:"api_key_name,omitempty"`
}
