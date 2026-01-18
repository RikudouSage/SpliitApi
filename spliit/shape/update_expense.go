package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

type UpdateExpenseRequest struct {
	ExpenseID     string            `json:"expenseId"`
	GroupID       string            `json:"groupId"`
	FormValues    ModifyExpenseForm `json:"expenseFormValues"`
	ParticipantID *string           `json:"participantId,omitempty"`
}

// todo consolidate
func (receiver *UpdateExpenseRequest) ApplyDefaults() {
	if receiver.FormValues.SplitMode == "" {
		receiver.FormValues.SplitMode = model.SplitModeEvenly
	}
	if receiver.FormValues.RecurrenceRule == "" {
		receiver.FormValues.RecurrenceRule = model.RecurrenceRuleNone
	}
}

type UpdateExpenseResponse CreateExpenseResponse
