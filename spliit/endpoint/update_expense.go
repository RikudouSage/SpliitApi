package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type UpdateExpense struct{}

func (receiver *UpdateExpense) Name() string {
	return "groups.expenses.update"
}

func (receiver *UpdateExpense) InputShape() shape.UpdateExpenseRequest {
	return shape.UpdateExpenseRequest{}
}

func (receiver *UpdateExpense) OutputShape() shape.UpdateExpenseResponse {
	return shape.UpdateExpenseResponse{}
}

func (receiver *UpdateExpense) Mutates() bool {
	return true
}
