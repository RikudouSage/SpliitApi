package model

import "time"

type ActivityType string

const (
	ActivityTypeUpdateGroup   ActivityType = "UPDATE_GROUP"
	ActivityTypeCreateExpense ActivityType = "CREATE_EXPENSE"
	ActivityTypeUpdateExpense ActivityType = "UPDATE_EXPENSE"
	ActivityTypeDeleteExpense ActivityType = "DELETE_EXPENSE"
)

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
