package model

import "time"

type Attendee struct {
	MemberID    uint64
	GatheringID uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
