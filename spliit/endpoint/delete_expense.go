package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type DeleteExpense struct {
}

func (receiver *DeleteExpense) Name() string {
	return "groups.expenses.delete"
}

func (receiver *DeleteExpense) InputShape() shape.DeleteExpenseRequest {
	return shape.DeleteExpenseRequest{}
}

func (receiver *DeleteExpense) OutputShape() shape.DeleteExpenseResponse {
	return shape.DeleteExpenseResponse{}
}

func (receiver *DeleteExpense) Mutates() bool {
	return true
}
