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

func (p *PgRepository) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	tx, err := p.dbCore.BeginTx(ctx, opts)
	if err != nil {
		return &sql.Tx{}, err
	}
	return tx, err
}

func (p *PgRepository) CommitTx(tx *sql.Tx) error {
	return tx.Commit()
}

func (p *PgRepository) RollbackTx(tx *sql.Tx) error {
	return tx.Rollback()
}

type Repository interface {
	IsEmailExists(ctx context.Context, email string) (bool, error)
	StoreMember(ctx context.Context, regData model.RegistrationRequest) error
	GetMemberByFilter(ctx context.Context, filter map[string]string) ([]model.Member, error)
	GetGatheringByFilter(ctx context.Context, filter map[string]string) ([]model.Gathering, error)
	GetAttendeeByFilter(ctx context.Context, filter map[string]string) ([]model.Attendee, error)
	StoreAttendee(ctx context.Context, tx *sql.Tx, memberID, gatheringID uint64) error
	UpdateInvStatus(ctx context.Context, tx *sql.Tx, memberID, gatheringID uint64, status int) error
	GetGatheringTypeByFilter(ctx context.Context, filter map[string]string) ([]model.GatheringType, error)
	StoreGathering(ctx context.Context, tx *sql.Tx, g model.Gathering) error
	StoreBulkInvitation(ctx context.Context, tx *sql.Tx, gatheringID uint64, invID []uint64) error

	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	CommitTx(tx *sql.Tx) error
	RollbackTx(tx *sql.Tx) error
}
