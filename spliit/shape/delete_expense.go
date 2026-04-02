package shape

// DeleteExpenseRequest is the input for the delete expense endpoint.
type DeleteExpenseRequest struct {
	ExpenseID     string  `json:"expenseId"`
	GroupID       string  `json:"groupId"`
	ParticipantID *string `json:"participantId,omitempty"`
}

// DeleteExpenseResponse is returned by the delete expense endpoint.
type DeleteExpenseResponse struct{}
