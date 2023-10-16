package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/azinudinachzab/hukumonline/model"
)

func (p *PgRepository) IsEmailExists(ctx context.Context, email string) (bool, error) {
	q := `SELECT email FROM users WHERE email = ?;`

	var emailDB string
	err := p.dbCore.QueryRowContext(ctx, q, email).Scan(&emailDB)
	if errors.Is(err, sql.ErrNoRows) {
		return false, model.ErrNotFound
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *PgRepository) IsPhoneNumberExists(ctx context.Context, phone string) (bool, error) {
	q := `SELECT phone_number FROM users WHERE phone_number = ?;`

	var phoneDB string
	err := p.dbCore.QueryRowContext(ctx, q, phone).Scan(&phoneDB)
	if errors.Is(err, sql.ErrNoRows) {
		return false, model.ErrNotFound
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *PgRepository) StoreUser(ctx context.Context, regData model.RegistrationRequest) error {
	q := `INSERT INTO users (phone_number, email, id_no, birth_place, birth_date, full_name, gender, verification_status)
			VALUES (?,?,?,?,?,?,?,?);`

	if _, err := p.dbCore.ExecContext(ctx, q, &regData.PhoneNumber, &regData.Email, &regData.IDNo, &regData.BirthPlace,
		&regData.BirthDate, &regData.FullName, &regData.Gender, &regData.VerificationStatus); err != nil {
		return err
	}

	return nil
}

func (p *PgRepository) GetUsersByFilter(ctx context.Context, filter map[string]string) ([]model.User, error) {
	query := `SELECT id, email, full_name, gender FROM users`
	if len(filter) > 0 {
		query += ` WHERE `
	}

	args := make([]interface{}, 0)
	idx := 1
	for key, val := range filter {
		query += fmt.Sprintf("%v = ?", key)

		if idx != len(filter) {
			query += " AND "
		}
		args = append(args, val)
		idx += 1
	}

	rows, err := p.dbCore.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userData := make([]model.User, 0)
	for rows.Next() {
		var (
			id            uint64
			email,fn,g    sql.NullString
		)
		if err := rows.Scan(&id, &email, &fn, &g); err != nil {
			return nil, err
		}

		userData = append(userData, model.User{
			ID:                 id,
			Email:              email.String,
			FullName:           fn.String,
			Gender:             g.String,
		})
	}

	return userData, nil
}
