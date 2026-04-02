package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// ListExpenses describes the groups.expenses.list endpoint.
type ListExpenses struct{}

// Name returns the remote endpoint name.
func (receiver *ListExpenses) Name() string {
	return "groups.expenses.list"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *ListExpenses) InputShape() shape.ListExpensesRequest {
	return shape.ListExpensesRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *ListExpenses) OutputShape() shape.ListExpensesResponse {
	return shape.ListExpensesResponse{}
}

// Mutates reports that the endpoint is read-only.
func (receiver *ListExpenses) Mutates() bool {
	return false
}
