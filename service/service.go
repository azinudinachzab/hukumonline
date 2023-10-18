package service

import (
	"context"
	"github.com/azinudinachzab/hukumonline/model"
	"github.com/azinudinachzab/hukumonline/repository"
	"github.com/go-playground/validator/v10"
)

type Dependency struct {
	Validator *validator.Validate
	Repo      repository.Repository
	Conf      model.Configuration
}

type AppService struct {
	validator *validator.Validate
	repo      repository.Repository
	conf      model.Configuration
}

func NewService(dep Dependency) *AppService {
	return &AppService{
		validator: dep.Validator,
		repo:      dep.Repo,
		conf:      dep.Conf,
	}
}

type Service interface {
	Registration(ctx context.Context, req model.RegistrationRequest) error
	CreateGathering(ctx context.Context, req model.CreateGatheringRequest) error
	ResponseInvitation(ctx context.Context, req model.ResponseInvitationRequest) error
}
