package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

// ListExpensesRequest is the input for the list expenses endpoint.
type ListExpensesRequest struct {
	GroupID string  `json:"groupId"`
	Cursor  *int    `json:"cursor,omitempty"`
	Limit   *int    `json:"limit,omitempty"`
	Filter  *string `json:"filter,omitempty"`
}

// ListExpensesResponse contains a page of expenses.
type ListExpensesResponse struct {
	Expenses   []model.Expense `json:"expenses"`
	HasMore    bool            `json:"hasMore"`
	NextCursor int             `json:"nextCursor"`
}
