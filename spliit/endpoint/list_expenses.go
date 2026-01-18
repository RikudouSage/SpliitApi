package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type ListExpenses struct{}

func (receiver *ListExpenses) Name() string {
	return "groups.expenses.list"
}

func (receiver *ListExpenses) InputShape() shape.ListExpensesRequest {
	return shape.ListExpensesRequest{}
}

func (receiver *ListExpenses) OutputShape() shape.ListExpensesResponse {
	return shape.ListExpensesResponse{}
}

func (receiver *ListExpenses) Mutates() bool {
	return false
}
