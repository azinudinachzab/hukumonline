package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/azinudinachzab/hukumonline/model"
)

func (p *PgRepository) IsEmailExists(ctx context.Context, email string) (bool, error) {
	q := `SELECT email FROM member WHERE email = ?;`

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

func (p *PgRepository) StoreMember(ctx context.Context, regData model.RegistrationRequest) error {
	q := `INSERT INTO member (first_name, last_name, email) VALUES (?,?,?);`

	if _, err := p.dbCore.ExecContext(ctx, q, &regData.FirstName, &regData.LastName, &regData.Email); err != nil {
		return err
	}

	return nil
}

func (p *PgRepository) GetMemberByFilter(ctx context.Context, filter map[string]string) ([]model.Member, error) {
	query := `SELECT id, email, first_name, last_name FROM member`
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

	memberData := make([]model.Member, 0)
	for rows.Next() {
		var (
			id            uint64
			email,fn,ln    sql.NullString
		)
		if err := rows.Scan(&id, &email, &fn, &ln); err != nil {
			return nil, err
		}

		memberData = append(memberData, model.Member{
			ID:        id,
			FirstName: fn.String,
			LastName:  ln.String,
			Email:     email.String,
		})
	}

	return memberData, nil
}
