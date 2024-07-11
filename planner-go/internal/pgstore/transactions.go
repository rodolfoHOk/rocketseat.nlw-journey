package pgstore

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rodolfoHOk/rocketseat.nlw-journey/planner-go/internal/api/spec"
)

func (q *Queries) CreateTrip(ctx context.Context, pool *pgxpool.Pool, params spec.CreateTripRequest) (uuid.UUID, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("pgstore: failed to begin tx for create trip: %w", err)
	}
	defer tx.Rollback(ctx)

	qtx := q.WithTx(tx)
	tripID, err := qtx.InsertTrip(ctx, InsertTripParams{
		Destination: params.Destination,
		OwnerEmail:  string(params.OwnerEmail),
		OwnerName:   params.OwnerName,
		StartsAt:    pgtype.Timestamp{Valid: true, Time: params.StartsAt},
		EndsAt:      pgtype.Timestamp{Valid: true, Time: params.EndsAt},
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("pgstore: failed to insert trip for create trip: %w", err)
	}

	params.EmailsToInvite = append(params.EmailsToInvite, params.OwnerEmail)
	participants := make([]InviteParticipantsToTripParams, len(params.EmailsToInvite))
	for i, eti := range params.EmailsToInvite {
		participants[i] = InviteParticipantsToTripParams{
			TripID: tripID,
			Email:  string(eti),
		}
	}
	if _, err := qtx.InviteParticipantsToTrip(ctx, participants); err != nil {
		return uuid.UUID{}, fmt.Errorf("pgstore: failed to insert participants for create trip: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.UUID{}, fmt.Errorf("pgstore: failed to commit tx for create trip: %w", err)
	}

	return tripID, nil
}
