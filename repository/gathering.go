package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/azinudinachzab/hukumonline/model"
)

func (p *PgRepository) GetGatheringByFilter(ctx context.Context, filter map[string]string) ([]model.Gathering, error) {
	query := `SELECT id, name, location, scheduled_at FROM gathering`
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

	gatheringData := make([]model.Gathering, 0)
	for rows.Next() {
		var (
			id          uint64
			name,loc    sql.NullString
			sched       sql.NullTime
		)
		if err := rows.Scan(&id, &name, &loc, &sched); err != nil {
			return nil, err
		}

		gatheringData = append(gatheringData, model.Gathering{
			ID:          id,
			ScheduledAt: sched.Time,
			Name:        name.String,
			Location:    loc.String,
		})
	}

	return gatheringData, nil
}

func (p *PgRepository) StoreGathering(ctx context.Context, tx *sql.Tx, g model.Gathering) error {
	dbExec := p.dbCore.ExecContext
	if tx != nil {
		dbExec = tx.ExecContext
	}
	q := `INSERT INTO gathering (id, creator, type, scheduled_at, name, location) VALUES (?,?,?,?,?,?);`

	if _, err := dbExec(ctx, q, &g.ID, &g.Creator, &g.Type, &g.ScheduledAt, &g.Name, &g.Location); err != nil {
		return err
	}

	return nil
}
