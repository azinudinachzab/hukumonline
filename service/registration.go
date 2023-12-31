package service

import (
	"context"
	"errors"
	"log"

	"github.com/azinudinachzab/hukumonline/model"
	"github.com/azinudinachzab/hukumonline/pkg/errs"
)

func (s *AppService) Registration(ctx context.Context, req model.RegistrationRequest) error {
	if err := s.validator.Struct(req); err != nil {
		return errs.NewWithErr(model.ECodeValidateFail, "validation request failed", err)
	}

	okEmail, err := s.repo.IsEmailExists(ctx, req.Email)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return errs.New(model.ECodeInternal, "failed to check email")
	}

	if okEmail {
		log.Println("email is exists")
		return errs.NewWithAttribute(model.ECodeDataExists, "email is exists", []errs.Attribute{{
			Field:   "email",
			Message: "email is exists",
		}})
	}

	err = s.repo.StoreMember(ctx, req)
	if err != nil {
		log.Println("failed to store user")
		return errs.New(model.ECodeInternal, "failed to store user")
	}

	return nil
}
