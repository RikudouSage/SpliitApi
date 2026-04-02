package shape

import (
	"go.chrastecky.dev/spliit-api/spliit/amount"

	"github.com/shopspring/decimal"
)

// GetStatsRequest is the input for the stats endpoint.
type GetStatsRequest struct {
	GroupID       string  `json:"groupId"`
	ParticipantID *string `json:"participantId,omitempty"`
}

// GetStatsResponse contains aggregated spend statistics.
type GetStatsResponse struct {
	TotalGroupSpendings       amount.Amount    `json:"totalGroupSpendings"`
	TotalParticipantSpendings *amount.Amount   `json:"totalParticipantSpendings"`
	TotalParticipantShare     *decimal.Decimal `json:"totalParticipantShare"`
}
