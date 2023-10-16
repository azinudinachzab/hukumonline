package model

import (
	"time"
)

type (
	User struct {
		ID                  uint64
		PhoneNumber         string
		Email               string
		IDNo                string
		BirthPlace          string
		BirthDate           time.Time
		FullName            string
		Gender              string
		VerificationStatus  int
		CreatedAt           time.Time
		UpdatedAt           time.Time
	}

	RegistrationRequest struct {
		PhoneNumber         string    `validate:"required,min=10,max=15,omitempty" json:"phone_number"`
		Email               string    `validate:"required,email,omitempty" json:"email"`
		IDNo                string    `validate:"required,omitempty" json:"id_no"`
		BirthPlace          string    `validate:"required,omitempty" json:"birth_place"`
		BirthDate           string    `validate:"required,omitempty" json:"birth_date"`
		FullName            string    `validate:"required,omitempty" json:"full_name"`
		Gender              string    `validate:"required,oneof=L P,omitempty" json:"gender"`
		VerificationStatus  int       `validate:"oneof=0 1,omitempty" json:"verification_status"`
	}
)

