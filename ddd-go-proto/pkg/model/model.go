package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type AuthenticationClient struct {
	AuthenticationClientType AuthenticationClientType `json:"authentication_client_type" validate:"required"`
	ClientId                 string                   `json:"client_id,omitempty"`
	PlatformAccountId        int64                    `json:"platform_account_id,omitempty" validate:"required"`
	ClientSecret             string                   `json:"client_secret,omitempty"`
	ApplicationUuid          string                   `json:"application_uuid,omitempty" validate:"required,gte=36,lte=36"`
	Enabled                  bool                     `json:"enabled,omitempty" validate:"required"`
	UserId                   int64                    `json:"user_id,omitempty"`
	ApiKey                   string                   `json:"api_key,omitempty"`
	ApiKeyName               string                   `json:"api_key_name,omitempty"`
	CreatedAt                time.Time                `json:"created_at,omitempty"`
	UpdatedAt                time.Time                `json:"updated_at,omitempty"`
}

func (m AuthenticationClient) Validate(validate *validator.Validate) *Error {
	var errors []ValidationError
	err := validate.Struct(&m)
	if gvError, ok := err.(validator.ValidationErrors); ok {
		for i := 0; i < len(gvError); i++ {
			validationError := ToValidationError(gvError[i])
			errors = append(errors, validationError)
		}
	}

	if len(errors) != 0 {
		// return error don't continue.
		return NewError("Validations", "invalid request", errors)
	}

	switch m.AuthenticationClientType {
	case API_KEY:
		values := map[string]interface{}{"client_id": m.ClientId, "client_secret": m.ClientSecret, "api_key_name": m.ApiKeyName, "user_id": m.UserId}
		rules := map[string]interface{}{"client_id": "required,gte=1,lte=36", "client_secret": "required,gte=1,lte=256", "api_key_name": "required,gte=1,lte=128", "user_id": "required,ne=0"}

		validateMap := validate.ValidateMap(values, rules)
		for fieldName, err := range validateMap {
			if gvError, ok := err.(validator.ValidationErrors); ok {
				for i := 0; i < len(gvError); i++ {
					validationError := ToValidationErrorWithName(fieldName, gvError[i])
					errors = append(errors, validationError)
				}
			}
		}
		if len(errors) != 0 {
			return NewError("Validations", "invalid request", errors)
		}
		return nil

	case PRIVATE_TOKEN, PUBLIC_TOKEN:
		// todo no validation for api_key
		values := map[string]interface{}{"application_uuid": m.ApplicationUuid}
		rules := map[string]interface{}{"application_uuid": "required,gte=36,lte=36"}

		validateMap := validate.ValidateMap(values, rules)
		for fieldName, err := range validateMap {
			if gvError, ok := err.(validator.ValidationErrors); ok {
				for i := 0; i < len(gvError); i++ {
					validationError := ToValidationErrorWithName(fieldName, gvError[i])
					errors = append(errors, validationError)
				}
			}
		}
		if len(errors) != 0 {
			return NewError("Validations", "invalid request", errors)
		}
		return nil
	case JWT:
		values := map[string]interface{}{"client_id": m.ClientId, "client_secret": m.ClientSecret}
		rules := map[string]interface{}{"client_id": "required,gte=1,lte=36", "client_secret": "required,gte=1,lte=256"}

		validateMap := validate.ValidateMap(values, rules)
		for fieldName, err := range validateMap {
			if gvError, ok := err.(validator.ValidationErrors); ok {
				for i := 0; i < len(gvError); i++ {
					validationError := ToValidationErrorWithName(fieldName, gvError[i])
					errors = append(errors, validationError)
				}
			}
		}
		if len(errors) != 0 {
			return NewError("Validations", "invalid request", errors)
		}
		return nil
	default:
		err := NewValidationError("authentication_client_type", "invalid authentication_client_type", InvalidRestriction)
		return NewError("Validations", "invalid request", []ValidationError{err})
	}
}

func ToValidationError(fieldError validator.FieldError) ValidationError {
	var msg string
	if fieldError.Param() != "" {
		msg = fmt.Sprintf("validation failed on '%s=%s' tag", fieldError.Tag(), fieldError.Param())
	} else {
		msg = fmt.Sprintf("validation failed on '%s' tag", fieldError.Tag())
	}
	validationError := NewValidationError(fieldError.Field(), msg, RestrictionType(fieldError.Tag()))
	return validationError
}

func ToValidationErrorWithName(fieldName string, fieldError validator.FieldError) ValidationError {
	var msg string
	if fieldError.Param() != "" {
		msg = fmt.Sprintf("validation failed on '%s=%s' tag", fieldError.Tag(), fieldError.Param())
	} else {
		msg = fmt.Sprintf("validation failed on '%s' tag", fieldError.Tag())
	}
	validationError := NewValidationError(fieldName, msg, RestrictionType(fieldError.Tag()))
	return validationError
}
