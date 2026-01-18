package shape

import (
	"go.chrastecky.dev/spliit-api/spliit/amount"

	"github.com/shopspring/decimal"
)

type GetStatsRequest struct {
	GroupID       string  `json:"groupId"`
	ParticipantID *string `json:"participantId,omitempty"`
}

type GetStatsResponse struct {
	TotalGroupSpendings       amount.Amount    `json:"totalGroupSpendings"`
	TotalParticipantSpendings *amount.Amount   `json:"totalParticipantSpendings"`
	TotalParticipantShare     *decimal.Decimal `json:"totalParticipantShare"`
}
