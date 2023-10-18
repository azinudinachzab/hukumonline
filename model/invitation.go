package model

type (
	ResponseInvitationRequest struct {
		MemberID    uint64 `validate:"required,gt=0"`
		GatheringID uint64 `validate:"required,gt=0"`
		Response    int    `validate:"required,oneof=1 2"`
	}
)
