package repository

import (
	"context"
	"database/sql"
	"fmt"
)

func (p *PgRepository) UpdateInvStatus(ctx context.Context, tx *sql.Tx, memberID, gatheringID uint64, status int) error {
	dbExec := p.dbCore.ExecContext
	if tx != nil {
		dbExec = tx.ExecContext
	}
	q := `UPDATE invitation SET status = ? WHERE member_id = ? AND gathering_id = ?`

	if _, err := dbExec(ctx, q, &status, &memberID, &gatheringID); err != nil {
		return err
	}

	return nil
}

func (p *PgRepository) StoreBulkInvitation(ctx context.Context, tx *sql.Tx, gatheringID uint64, invID []uint64) error {
	dbExec := p.dbCore.ExecContext
	if tx != nil {
		dbExec = tx.ExecContext
	}
	valueString := ""
	args := make([]interface{}, 0)
	for i, val := range invID {
		valueString += "("
		args = append(args, val)
		valueString += fmt.Sprintf("%s,", "?")
		args = append(args, gatheringID)
		valueString += fmt.Sprintf("%s", "?")
		valueString += ")"

		if i == len(invID) - 1 {
			valueString += ";"
		} else {
			valueString += ","
		}
	}

	q := `INSERT INTO invitation (member_id, gathering_id) VALUES ` + valueString

	if _, err := dbExec(ctx, q, args...); err != nil {
		return err
	}

	return nil
}
