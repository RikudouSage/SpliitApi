package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

// UpdateExpenseRequest is the input for the update expense endpoint.
type UpdateExpenseRequest struct {
	ExpenseID     string            `json:"expenseId"`
	GroupID       string            `json:"groupId"`
	FormValues    ModifyExpenseForm `json:"expenseFormValues"`
	ParticipantID *string           `json:"participantId,omitempty"`
}

// ApplyDefaults applies default split and recurrence values expected by the API.
func (receiver *UpdateExpenseRequest) ApplyDefaults() {
	if receiver.FormValues.SplitMode == "" {
		receiver.FormValues.SplitMode = model.SplitModeEvenly
	}
	if receiver.FormValues.RecurrenceRule == "" {
		receiver.FormValues.RecurrenceRule = model.RecurrenceRuleNone
	}
}

// UpdateExpenseResponse is the result returned after updating an expense.
type UpdateExpenseResponse CreateExpenseResponse
