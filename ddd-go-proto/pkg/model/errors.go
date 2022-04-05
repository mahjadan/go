package model

import (
	"fmt"
	"strings"
)

var ErrNotFound = Error{ErrorCode: "NOT_FOUND", Message: "recurso não encontrado", ValidationErrors: nil}

var ErrConflict = Error{ErrorCode: "CONFLICT", Message: "recurso já existe", ValidationErrors: nil}

var ErrInternal = Error{ErrorCode: "INTERNAL", Message: "erro interno aconteceu", ValidationErrors: nil}

type Error struct {
	ErrorCode        string            `json:"error_code,omitempty"`
	Message          string            `json:"message"`
	ValidationErrors []ValidationError `json:"validation_errors,omitempty"`
}

func (er Error) Error() string {
	return er.String()
}

type ValidationError struct {
	Field       string          `json:"field"`
	Restriction RestrictionType `json:"restriction"`
	Message     string          `json:"message"`
}

type RestrictionType string

// this can be used to manually create validationError
// it could be adapted to use same patter of playground validation tag names example: 'eq' , 'gte'
// or we can map the values 'gte' to 'MIN_OR_EQUAL' and 'gt' to 'MIN'
const (
	RequiredRestriction       RestrictionType = "REQUIRED"
	MatchRestriction          RestrictionType = "MATCH"
	InvalidRestriction        RestrictionType = "INVALID"
	MinLengthRestriction      RestrictionType = "MIN_LENGTH"
	MaxLengthRestriction      RestrictionType = "MAX_LENGTH"
	MismatchLengthRestriction RestrictionType = "MISMATCH_LENGTH"
	NumericValueRestriction   RestrictionType = "NUMERIC_VALUE"
	ExactLengthRestriction    RestrictionType = "EXACT_LENGTH"
	GreaterThanRestriction    RestrictionType = "GREATER_THAN"
	LessThanRestriction       RestrictionType = "LESS_THAN"
)

func NewError(errorCode, msg string, validationErrors []ValidationError) *Error {
	return &Error{
		ErrorCode:        errorCode,
		Message:          msg,
		ValidationErrors: validationErrors,
	}
}
func NewValidationError(fieldName, message string, restrictionType RestrictionType) ValidationError {
	return ValidationError{
		Field:       fieldName,
		Restriction: restrictionType,
		Message:     message,
	}
}

func NewInternalError(message string) *Error {
	return &Error{ErrorCode: "INTERNAL", Message: fmt.Sprintf("erro interno aconteceu, %s", message), ValidationErrors: nil}
}
func (v ValidationError) Error() string {
	return v.String()
}
func (er *Error) String() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("errorCode:%s, ", er.ErrorCode))
	s.WriteString(fmt.Sprintf("message:%s, ", er.Message))
	for _, validationError := range er.ValidationErrors {
		s.WriteString(fmt.Sprintf("%s, ", validationError.String()))
	}
	return s.String()
}
func (v ValidationError) String() string {
	return fmt.Sprintf("validationError: %s:%s:%s", v.Field, v.Restriction, v.Message)
}
