package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// CreateExpense describes the groups.expenses.create endpoint.
type CreateExpense struct {
}

// Name returns the remote endpoint name.
func (receiver *CreateExpense) Name() string {
	return "groups.expenses.create"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *CreateExpense) InputShape() shape.CreateExpenseRequest {
	return shape.CreateExpenseRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *CreateExpense) OutputShape() shape.CreateExpenseResponse {
	return shape.CreateExpenseResponse{}
}

// Mutates reports that the endpoint changes remote state.
func (receiver *CreateExpense) Mutates() bool {
	return true
}
