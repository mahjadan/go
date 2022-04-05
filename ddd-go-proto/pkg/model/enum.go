package model

import (
	"encoding/json"
)

type AuthenticationClientType int

const (
	UNKNOWN AuthenticationClientType = iota
	JWT
	PUBLIC_TOKEN
	PRIVATE_TOKEN
	API_KEY
)

func (t AuthenticationClientType) String() string {
	return [...]string{"UNKNOWN", "JWT", "PUBLIC_TOKEN", "PRIVATE_TOKEN", "API_KEY"}[t]
}

func (t *AuthenticationClientType) FromString(authenticationClientType string) AuthenticationClientType {
	return map[string]AuthenticationClientType{
		"UNKNOWN":       UNKNOWN,
		"JWT":           JWT,
		"PUBLIC_TOKEN":  PUBLIC_TOKEN,
		"PRIVATE_TOKEN": PRIVATE_TOKEN,
		"API_KEY":       API_KEY,
	}[authenticationClientType]
}

func (t AuthenticationClientType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *AuthenticationClientType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = t.FromString(s)
	return nil
}
