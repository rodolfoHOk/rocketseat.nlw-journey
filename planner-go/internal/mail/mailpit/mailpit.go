package mailpit

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rodolfoHOk/rocketseat.nlw-journey/planner-go/internal/pgstore"
	"github.com/wneessen/go-mail"
)

type store interface {
	GetTrip(context.Context, uuid.UUID) (pgstore.GetTripRow, error)
}

type Mailpit struct {
	store store
}

func NewMailPit(pool *pgxpool.Pool) Mailpit {
	return Mailpit{
		pgstore.New(pool),
	}
}

func (mp Mailpit) SendConfirmTripEmailToTripOwner(tripID uuid.UUID) error {
	ctx := context.Background()
	trip, err := mp.store.GetTrip(ctx, tripID)
	if err != nil {
		return fmt.Errorf("mailpit: failed to get trip for SendConfirmTripEmailToTripOwner: %w", err)
	}

	msg := mail.NewMsg()
	if err := msg.From("mailpit@journey.com"); err != nil {
		return fmt.Errorf("mailpit: failed to set 'from' in email SendConfirmTripEmailToTripOwner: %w", err)
	}
	if err := msg.To(trip.OwnerEmail); err != nil {
		return fmt.Errorf("mailpit: failed to set 'to' in email SendConfirmTripEmailToTripOwner: %w", err)
	}
	msg.Subject("Confirme sua viagem")
	msg.SetBodyString(mail.TypeTextPlain, fmt.Sprintf(`
		Olá, %s!

		A sua viagem para %s que começa no dia %s precisa ser confirmada.
		Clique no botão abaixo para confirmar.
		`,
		trip.OwnerName, trip.Destination, trip.StartsAt.Time.Format(time.DateOnly),
	))

	client, err := mail.NewClient("localhost", mail.WithTLSPortPolicy(mail.NoTLS), mail.WithPort(1025))
	if err != nil {
		return fmt.Errorf("mailpit: failed create email client SendConfirmTripEmailToTripOwner: %w", err)
	}

	if err := client.DialAndSend(msg); err != nil {
		return fmt.Errorf("mailpit: failed send email client SendConfirmTripEmailToTripOwner: %w", err)
	}

	return nil
}