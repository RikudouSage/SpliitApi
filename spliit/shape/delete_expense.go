package shape

type DeleteExpenseRequest struct {
	ExpenseID     string  `json:"expenseId"`
	GroupID       string  `json:"groupId"`
	ParticipantID *string `json:"participantId,omitempty"`
}

type DeleteExpenseResponse struct{}
