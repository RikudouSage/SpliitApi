package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// GetExpense describes the groups.expenses.get endpoint.
type GetExpense struct {
}

// Name returns the remote endpoint name.
func (receiver *GetExpense) Name() string {
	return "groups.expenses.get"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *GetExpense) InputShape() shape.GetExpenseRequest {
	return shape.GetExpenseRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *GetExpense) OutputShape() shape.GetExpenseResponse {
	return shape.GetExpenseResponse{}
}

// Mutates reports that the endpoint is read-only.
func (receiver *GetExpense) Mutates() bool {
	return false
}
