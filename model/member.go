package model

import (
	"time"
)

type (
	Member struct {
		ID        uint64
		FirstName string
		LastName  string
		Email     string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	RegistrationRequest struct {
		FirstName string `validate:"required,omitempty" json:"first_name"`
		LastName  string `validate:"required,omitempty" json:"last_name"`
		Email     string `validate:"required,email,omitempty" json:"email"`
	}
)
