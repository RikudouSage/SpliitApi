package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type GetExpense struct {
}

func (receiver *GetExpense) Name() string {
	return "groups.expenses.get"
}

func (receiver *GetExpense) InputShape() shape.GetExpenseRequest {
	return shape.GetExpenseRequest{}
}

func (receiver *GetExpense) OutputShape() shape.GetExpenseResponse {
	return shape.GetExpenseResponse{}
}

func (receiver *GetExpense) Mutates() bool {
	return false
}
