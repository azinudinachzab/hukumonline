package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/azinudinachzab/hukumonline/model"
)

func (p *PgRepository) GetGatheringTypeByFilter(ctx context.Context, filter map[string]string) (
	[]model.GatheringType, error) {
	query := `SELECT id, name FROM gathering_type`
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

	gatheringTypeData := make([]model.GatheringType, 0)
	for rows.Next() {
		var (
			id     uint64
			name   sql.NullString
		)
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		gatheringTypeData = append(gatheringTypeData, model.GatheringType{
			ID:          id,
			Name:        name.String,
		})
	}

	return gatheringTypeData, nil
}
