package shape

import "go.chrastecky.dev/spliit-api/spliit/amount"

type ListBalancesRequest struct {
	GroupID string `json:"groupId"`
}

type Reimbursement struct {
	From   string        `json:"from"`
	To     string        `json:"to"`
	Amount amount.Amount `json:"amount"`
}

type Balance struct {
	Paid    amount.Amount `json:"paid"`
	PaidFor amount.Amount `json:"paidFor"`
	Total   amount.Amount `json:"total"`
}

type ListBalancesResponse struct {
	Balances       map[string]Balance `json:"balances"`
	Reimbursements []Reimbursement    `json:"reimbursements"`
}
