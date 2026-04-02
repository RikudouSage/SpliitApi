package model

import "time"

// ActivityType identifies the kind of group activity entry.
type ActivityType string

const (
	// ActivityTypeUpdateGroup marks a group update event.
	ActivityTypeUpdateGroup ActivityType = "UPDATE_GROUP"
	// ActivityTypeCreateExpense marks an expense creation event.
	ActivityTypeCreateExpense ActivityType = "CREATE_EXPENSE"
	// ActivityTypeUpdateExpense marks an expense update event.
	ActivityTypeUpdateExpense ActivityType = "UPDATE_EXPENSE"
	// ActivityTypeDeleteExpense marks an expense deletion event.
	ActivityTypeDeleteExpense ActivityType = "DELETE_EXPENSE"
)

// Activity represents a single activity feed entry.
type Activity struct {
	ID            string       `json:"id"`
	Group         *Group       `json:"group"`
	GroupID       string       `json:"groupId"`
	Time          time.Time    `json:"time"`
	ActivityType  ActivityType `json:"activityType"`
	ParticipantID *string      `json:"participantId"`
	Expense       *Expense     `json:"expense"`
	ExpenseID     *string      `json:"expenseId"`
	Data          *string      `json:"data"`
}
