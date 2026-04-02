package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// DeleteExpense describes the groups.expenses.delete endpoint.
type DeleteExpense struct {
}

// Name returns the remote endpoint name.
func (receiver *DeleteExpense) Name() string {
	return "groups.expenses.delete"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *DeleteExpense) InputShape() shape.DeleteExpenseRequest {
	return shape.DeleteExpenseRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *DeleteExpense) OutputShape() shape.DeleteExpenseResponse {
	return shape.DeleteExpenseResponse{}
}

// Mutates reports that the endpoint changes remote state.
func (receiver *DeleteExpense) Mutates() bool {
	return true
}
