package repository

import (
	"context"
	"database/sql"

	"github.com/azinudinachzab/hukumonline/model"
)

type PgRepository struct {
	dbCore *sql.DB
}

func NewPgRepository(dbCore *sql.DB) *PgRepository {
	return &PgRepository{
		dbCore: dbCore,
	}
}

type Repository interface {
	IsEmailExists(ctx context.Context, email string) (bool, error)
	IsPhoneNumberExists(ctx context.Context, phone string) (bool, error)
	StoreUser(ctx context.Context, regData model.RegistrationRequest) error
	GetUsersByFilter(ctx context.Context, filter map[string]string) ([]model.User, error)
	StoreBulkPromo(ctx context.Context, promos []model.Promo) error
}
