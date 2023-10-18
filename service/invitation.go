package service

import (
	"context"
	"strconv"
	"time"

	"github.com/azinudinachzab/hukumonline/model"
	"github.com/azinudinachzab/hukumonline/pkg/errs"
)

const (
	errResponseInvitation = "failed to response invitation"
	errCreateGathering    = "failed to create gathering"
)

func (s *AppService) CreateGathering(ctx context.Context, req model.CreateGatheringRequest) error {
	if err := s.validator.Struct(req); err != nil {
		return errs.NewWithErr(model.ECodeValidateFail, "validation request failed", err)
	}
	// check invitation request
	if len(req.InvitedMember) < 1 {
		return errs.NewWithErr(model.ECodeValidateFail, "validation request failed", model.ErrRequestNotValid)
	}

	scheduleTime, err := time.Parse("2006-01-02 15:04:05", req.ScheduledAt)
	if err != nil {
		return errs.NewWithErr(model.ECodeValidateFail, "validation request failed", err)
	}

	// check creator
	members, err := s.repo.GetMemberByFilter(ctx, map[string]string{
		"id": strconv.FormatUint(req.CreatorID, 10),
	})
	if err != nil {
		return errs.NewWithErr(model.ECodeInternal, "failed to get member", err)
	}

	if len(members) < 1 {
		return errs.NewWithErr(model.ECodeNotFound, "member not found", model.ErrNotFound)
	}

	// check gathering type
	types, err := s.repo.GetGatheringTypeByFilter(ctx, map[string]string{
		"id": strconv.Itoa(req.Type),
	})
	if err != nil {
		return errs.NewWithErr(model.ECodeInternal, "failed to get type", err)
	}
	if len(types) < 1 {
		return errs.NewWithErr(model.ECodeNotFound, "type not found", model.ErrNotFound)
	}

	// store to gathering and invitation
	tx, err := s.repo.BeginTx(ctx, nil)
	if err != nil {
		return errs.NewWithErr(model.ECodeInternal, errCreateGathering, err)
	}

	gID := uint64(time.Now().UnixNano())
	if err := s.repo.StoreGathering(ctx, tx, model.Gathering{
		ID:          gID,
		Creator:     req.CreatorID,
		Type:        req.Type,
		ScheduledAt: scheduleTime,
		Name:        req.Name,
		Location:    req.Location,
	}); err != nil {
		if errRollback := s.repo.RollbackTx(tx); err != nil {
			return errs.NewWithErr(model.ECodeInternal, errCreateGathering, errRollback)
		}
		return errs.NewWithErr(model.ECodeInternal, errCreateGathering, err)
	}

	if err := s.repo.StoreBulkInvitation(ctx, tx, gID, req.InvitedMember); err != nil {
		if errRollback := s.repo.RollbackTx(tx); err != nil {
			return errs.NewWithErr(model.ECodeInternal, errCreateGathering, errRollback)
		}
		return errs.NewWithErr(model.ECodeInternal, errCreateGathering, err)
	}

	if err := s.repo.CommitTx(tx); err != nil {
		if errRollback := s.repo.RollbackTx(tx); err != nil {
			return errs.NewWithErr(model.ECodeInternal, errCreateGathering, errRollback)
		}
		return errs.NewWithErr(model.ECodeInternal, errCreateGathering, err)
	}

	return nil
}

func (s *AppService) ResponseInvitation(ctx context.Context, req model.ResponseInvitationRequest) error {
	if err := s.validator.Struct(req); err != nil {
		return errs.NewWithErr(model.ECodeValidateFail, "validation request failed", err)
	}

	// check if member exist
	members, err := s.repo.GetMemberByFilter(ctx, map[string]string{
		"id": strconv.FormatUint(req.MemberID, 10),
	})
	if err != nil {
		return errs.NewWithErr(model.ECodeInternal, "failed to get member", err)
	}

	if len(members) < 1 {
		return errs.NewWithErr(model.ECodeNotFound, "member not found", model.ErrNotFound)
	}

	// check if gathering exist and valid
	gatherings, err := s.repo.GetGatheringByFilter(ctx, map[string]string{
		"id": strconv.FormatUint(req.GatheringID, 10),
	})
	if err != nil {
		return errs.NewWithErr(model.ECodeInternal, "failed to get gathering", err)
	}

	if len(gatherings) < 1 {
		return errs.NewWithErr(model.ECodeNotFound, "gathering not found", model.ErrNotFound)
	}
	gathering := gatherings[0]

	if time.Now().After(gathering.ScheduledAt) {
		return errs.NewWithErr(model.ECodeNotFound, "gathering is expired", model.ErrNotFound)
	}

	// check if member is already response
	atts, err := s.repo.GetAttendeeByFilter(ctx, map[string]string{
		"member_id": strconv.FormatUint(req.MemberID, 10),
		"gathering_id": strconv.FormatUint(req.GatheringID, 10),
	})
	if err != nil {
		return errs.NewWithErr(model.ECodeInternal, "failed to get member", err)
	}
	if len(atts) > 0 {
		return errs.NewWithErr(model.ECodeDataExists, "member already responded", model.ErrResourceAlreadyExists)
	}

	// record member response (update status on table invitation, add to table attendee, use tx)
	tx, err := s.repo.BeginTx(ctx, nil)
	if err != nil {
		return errs.NewWithErr(model.ECodeInternal, errResponseInvitation, err)
	}

	if err := s.repo.UpdateInvStatus(ctx, tx, req.MemberID, req.GatheringID, req.Response); err != nil {
		if errRollback := s.repo.RollbackTx(tx); err != nil {
			return errs.NewWithErr(model.ECodeInternal, errResponseInvitation, errRollback)
		}
		return errs.NewWithErr(model.ECodeInternal, errResponseInvitation, err)
	}

	if err := s.repo.StoreAttendee(ctx, tx, req.MemberID, req.GatheringID); err != nil {
		if errRollback := s.repo.RollbackTx(tx); err != nil {
			return errs.NewWithErr(model.ECodeInternal, errResponseInvitation, errRollback)
		}
		return errs.NewWithErr(model.ECodeInternal, errResponseInvitation, err)
	}

	if err := s.repo.CommitTx(tx); err != nil {
		if errRollback := s.repo.RollbackTx(tx); err != nil {
			return errs.NewWithErr(model.ECodeInternal, errResponseInvitation, errRollback)
		}
		return errs.NewWithErr(model.ECodeInternal, errResponseInvitation, err)
	}

	return nil
}
