package model

import "time"

type (
	Promo struct {
		ID                  uint64
		UserID              uint64
		PromoCode           string
		Status              int
		ExpiryAt            time.Time
		Percentage          float64
		CreatedAt           time.Time
		UpdatedAt           time.Time
	}
)
