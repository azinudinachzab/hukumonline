package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/azinudinachzab/hukumonline/model"
)

func (p *PgRepository) GetAttendeeByFilter(ctx context.Context, filter map[string]string) ([]model.Attendee, error) {
	query := `SELECT member_id, gathering_id FROM attendee`
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

	attData := make([]model.Attendee, 0)
	for rows.Next() {
		var (
			mID, gID uint64
		)
		if err := rows.Scan(&mID, &gID); err != nil {
			return nil, err
		}

		attData = append(attData, model.Attendee{
			MemberID:    mID,
			GatheringID: gID,
		})
	}

	return attData, nil
}

func (p *PgRepository) StoreAttendee(ctx context.Context, tx *sql.Tx, memberID, gatheringID uint64) error {
	dbExec := p.dbCore.ExecContext
	if tx != nil {
		dbExec = tx.ExecContext
	}
	q := `INSERT INTO attendee (member_id, gathering_id) VALUES (?,?);`

	if _, err := dbExec(ctx, q, &memberID, &gatheringID); err != nil {
		return err
	}

	return nil
}
