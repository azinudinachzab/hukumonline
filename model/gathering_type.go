package model

import "time"

type (
	GatheringType struct {
		ID          uint64
		Name        string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
