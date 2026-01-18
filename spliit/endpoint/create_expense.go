package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type CreateExpense struct {
}

func (receiver *CreateExpense) Name() string {
	return "groups.expenses.create"
}

func (receiver *CreateExpense) InputShape() shape.CreateExpenseRequest {
	return shape.CreateExpenseRequest{}
}

func (receiver *CreateExpense) OutputShape() shape.CreateExpenseResponse {
	return shape.CreateExpenseResponse{}
}

func (receiver *CreateExpense) Mutates() bool {
	return true
}
