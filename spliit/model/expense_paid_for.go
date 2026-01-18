package model

type ExpensePaidFor struct {
	Expense       *Expense     `json:"expense"`
	Participant   *Participant `json:"participant"`
	ExpenseID     string       `json:"expenseId"`
	ParticipantID string       `json:"participantId"`
	Shares        uint         `json:"shares"`
}
