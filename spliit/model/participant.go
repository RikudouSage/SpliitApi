package model

// Participant represents a member of a group.
type Participant struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Group           *Group            `json:"group"`
	GroupID         string            `json:"groupId"`
	ExpensesPaidBy  []*Expense        `json:"expensesPaidBy"`
	ExpensesPaidFor []*ExpensePaidFor `json:"expensesPaidFor"`
}
