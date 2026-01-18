package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

type GetExpenseRequest struct {
	GroupID   string `json:"groupId"`
	ExpenseID string `json:"expenseId"`
}

type GetExpenseResponse struct {
	Expense model.Expense `json:"expense"`
}
