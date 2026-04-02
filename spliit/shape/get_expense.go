package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

// GetExpenseRequest is the input for the get expense endpoint.
type GetExpenseRequest struct {
	GroupID   string `json:"groupId"`
	ExpenseID string `json:"expenseId"`
}

// GetExpenseResponse contains the requested expense.
type GetExpenseResponse struct {
	Expense model.Expense `json:"expense"`
}
