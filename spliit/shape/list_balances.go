package shape

import "go.chrastecky.dev/spliit-api/spliit/amount"

// ListBalancesRequest is the input for the list balances endpoint.
type ListBalancesRequest struct {
	GroupID string `json:"groupId"`
}

// Reimbursement describes a suggested settlement transfer.
type Reimbursement struct {
	From   string        `json:"from"`
	To     string        `json:"to"`
	Amount amount.Amount `json:"amount"`
}

// Balance contains one participant balance summary.
type Balance struct {
	Paid    amount.Amount `json:"paid"`
	PaidFor amount.Amount `json:"paidFor"`
	Total   amount.Amount `json:"total"`
}

// ListBalancesResponse contains balances and reimbursement suggestions.
type ListBalancesResponse struct {
	Balances       map[string]Balance `json:"balances"`
	Reimbursements []Reimbursement    `json:"reimbursements"`
}
