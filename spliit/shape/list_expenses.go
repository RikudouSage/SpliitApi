package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

type ListExpensesRequest struct {
	GroupID string  `json:"groupId"`
	Cursor  *int    `json:"cursor,omitempty"`
	Limit   *int    `json:"limit,omitempty"`
	Filter  *string `json:"filter,omitempty"`
}

type ListExpensesResponse struct {
	Expenses   []model.Expense `json:"expenses"`
	HasMore    bool            `json:"hasMore"`
	NextCursor int             `json:"nextCursor"`
}
