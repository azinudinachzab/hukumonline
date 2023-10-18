package model

import "time"

type (
	Gathering struct {
		ID          uint64
		Creator     uint64
		Type        int
		ScheduledAt time.Time
		Name        string
		Location    string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	CreateGatheringRequest struct {
		CreatorID     uint64   `validate:"required,gt=0" json:"creator_id"`
		Name          string   `validate:"required" json:"name"`
		Location      string   `validate:"required" json:"location"`
		ScheduledAt   string   `validate:"required" json:"scheduled_at"`
		Type          int      `validate:"required" json:"type"`
		InvitedMember []uint64 `json:"invited_member"`
	}
)
