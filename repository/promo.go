package repository

import (
	"context"
	"fmt"
	"github.com/azinudinachzab/hukumonline/model"
	"strings"
)

func (p *PgRepository) StoreBulkPromo(ctx context.Context, promos []model.Promo) error {
	valueStrings := make([]string, 0)
	valueArgs := make([]interface{}, 0)
	for _, p := range promos {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, p.UserID)
		valueArgs = append(valueArgs, p.PromoCode)
		valueArgs = append(valueArgs, p.Status)
		valueArgs = append(valueArgs, p.ExpiryAt)
		valueArgs = append(valueArgs, p.Percentage)
	}
	q := fmt.Sprintf(`INSERT INTO promos (user_id, code, status, expiry_at, percentage)
			VALUES %s`, strings.Join(valueStrings, ","))

	if _, err := p.dbCore.ExecContext(ctx, q, valueArgs...); err != nil {
		return err
	}

	return nil
}
