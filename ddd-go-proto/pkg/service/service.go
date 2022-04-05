package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go-proto-poc/pkg/model"
	"go-proto-poc/pkg/repository"
	"log"
)

type Service struct {
	repository repository.Repository
	validator  *validator.Validate
}

func (s Service) CreateAuthenticationClient(ctx context.Context, client model.AuthenticationClient) (model.AuthenticationClient, error) {
	err1 := client.Validate(s.validator)
	if err1 != nil {
		// we return our model.Error and the transport layer transform the error to whatever transport protocol it is using (json, proto)
		log.Println(err1)
		return model.AuthenticationClient{}, err1
	}
	authenticationClient, err := s.repository.CreateAuthentication(ctx, client)
	if err != nil {
		// in this example we could only have a 500 errors or Duplicated ID.
		return model.AuthenticationClient{}, model.NewInternalError(err.Error())
	}
	//transform the model into a repository model which will be persisted in the db
	// this model can change depending on the db as some databases require different annotations like mongodb ( 'bson' )
	return authenticationClient, nil
}

func NewService(validate *validator.Validate, repo repository.Repository) Service {
	return Service{
		repository: repo,
		validator:  validate,
	}
}
