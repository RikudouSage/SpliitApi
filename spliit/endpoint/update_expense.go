package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// UpdateExpense describes the groups.expenses.update endpoint.
type UpdateExpense struct{}

// Name returns the remote endpoint name.
func (receiver *UpdateExpense) Name() string {
	return "groups.expenses.update"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *UpdateExpense) InputShape() shape.UpdateExpenseRequest {
	return shape.UpdateExpenseRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *UpdateExpense) OutputShape() shape.UpdateExpenseResponse {
	return shape.UpdateExpenseResponse{}
}

// Mutates reports that the endpoint changes remote state.
func (receiver *UpdateExpense) Mutates() bool {
	return true
}
