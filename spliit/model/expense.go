package model

import (
	"go.chrastecky.dev/spliit-api/spliit/amount"
	"time"

	"github.com/shopspring/decimal"
)

type SplitMode string
type RecurrenceRule string

const (
	SplitModeEvenly        SplitMode = "EVENLY"
	SplitModeByShares      SplitMode = "BY_SHARES"
	SplitModeByPercentages SplitMode = "BY_PERCENTAGES"
	SplitModeByAmount      SplitMode = "BY_AMOUNT"

	RecurrenceRuleNone    RecurrenceRule = "NONE"
	RecurrenceRuleDaily   RecurrenceRule = "DAILY"
	RecurrenceRuleWeekly  RecurrenceRule = "WEEKLY"
	RecurrenceRuleMonthly RecurrenceRule = "MONTHLY"
)

type Expense struct {
	ID               string             `json:"id"`
	Group            *Group             `json:"group"`
	ExpenseDate      time.Time          `json:"expenseDate"`
	Title            string             `json:"title"`
	Category         *Category          `json:"category"`
	CategoryID       int                `json:"categoryId"`
	Amount           amount.Amount      `json:"amount"`
	OriginalAmount   *amount.Amount     `json:"originalAmount"`
	OriginalCurrency *string            `json:"originalCurrency"`
	ConversionRate   *decimal.Decimal   `json:"conversionRate"`
	PaidBy           *Participant       `json:"paidBy"`
	PaidById         string             `json:"paidById"`
	PaidFor          []*ExpensePaidFor  `json:"paidFor"`
	GroupID          string             `json:"groupId"`
	IsReimbursement  bool               `json:"isReimbursement"`
	SplitMode        SplitMode          `json:"splitMode"`
	CreatedAt        time.Time          `json:"createdAt"`
	Documents        []*ExpenseDocument `json:"documents"`
	Notes            *string            `json:"notes"`

	RecurrenceRule         *RecurrenceRule       `json:"recurrenceRule"`
	RecurringExpenseLink   *RecurringExpenseLink `json:"recurringExpenseLink"`
	RecurringExpenseLinkId *string               `json:"recurringExpenseLinkId"`

	// list expenses
	Counts map[string]uint `json:"_count"`
}
