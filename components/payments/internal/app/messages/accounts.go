package messages

import (
	"time"

	"github.com/formancehq/payments/internal/app/models"
	"github.com/formancehq/payments/pkg/events"
)

type accountMessagePayload struct {
	ID        string             `json:"id"`
	CreatedAt time.Time          `json:"createdAt"`
	Reference string             `json:"reference"`
	Provider  string             `json:"provider"`
	Type      models.AccountType `json:"type"`
}

func NewEventSavedAccounts(accounts []models.Account) events.EventMessage {
	payload := make([]accountMessagePayload, len(accounts))

	for accountIdx, account := range accounts {
		payload[accountIdx] = accountMessagePayload{
			ID:        account.ID.String(),
			CreatedAt: account.CreatedAt,
			Reference: account.Reference,
			Provider:  account.Provider,
			Type:      account.Type,
		}
	}

	return events.EventMessage{
		Date:    time.Now().UTC(),
		App:     events.EventApp,
		Version: events.EventVersion,
		Type:    events.EventTypeSavedAccounts,
		Payload: payload,
	}
}