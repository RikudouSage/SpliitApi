package model

import (
	"time"

	"go.chrastecky.dev/spliit-api/spliit/amount"

	"github.com/shopspring/decimal"
)

// SplitMode describes how an expense is split between participants.
type SplitMode string

// RecurrenceRule describes the repetition schedule of an expense.
type RecurrenceRule string

const (
	// SplitModeEvenly splits the expense equally between participants.
	SplitModeEvenly SplitMode = "EVENLY"
	// SplitModeByShares splits the expense according to share counts.
	SplitModeByShares SplitMode = "BY_SHARES"
	// SplitModeByPercentage splits the expense according to percentages.
	SplitModeByPercentage SplitMode = "BY_PERCENTAGE"
	// SplitModeByPercentages is retained as a source-compatible alias.
	//
	// Deprecated: use SplitModeByPercentage.
	SplitModeByPercentages = SplitModeByPercentage
	// SplitModeByAmount splits the expense according to explicit amounts.
	SplitModeByAmount SplitMode = "BY_AMOUNT"

	// RecurrenceRuleNone marks a non-recurring expense.
	RecurrenceRuleNone RecurrenceRule = "NONE"
	// RecurrenceRuleDaily marks a daily recurring expense.
	RecurrenceRuleDaily RecurrenceRule = "DAILY"
	// RecurrenceRuleWeekly marks a weekly recurring expense.
	RecurrenceRuleWeekly RecurrenceRule = "WEEKLY"
	// RecurrenceRuleMonthly marks a monthly recurring expense.
	RecurrenceRuleMonthly RecurrenceRule = "MONTHLY"
)

// Expense represents an expense stored in a Spliit group.
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
	PaidByID         string             `json:"paidById"`
	PaidFor          []*ExpensePaidFor  `json:"paidFor"`
	GroupID          string             `json:"groupId"`
	IsReimbursement  bool               `json:"isReimbursement"`
	SplitMode        SplitMode          `json:"splitMode"`
	CreatedAt        time.Time          `json:"createdAt"`
	Documents        []*ExpenseDocument `json:"documents"`
	Notes            *string            `json:"notes"`

	RecurrenceRule         *RecurrenceRule       `json:"recurrenceRule"`
	RecurringExpenseLink   *RecurringExpenseLink `json:"recurringExpenseLink"`
	RecurringExpenseLinkID *string               `json:"recurringExpenseLinkId"`

	// list expenses
	Counts map[string]uint `json:"_count"`
}
