package shape

import (
	"go.chrastecky.dev/spliit-api/spliit/amount"
	"go.chrastecky.dev/spliit-api/spliit/model"
	"time"

	"github.com/shopspring/decimal"
)

// ModifyExpenseFormPaidFor describes one participant share in an expense form.
type ModifyExpenseFormPaidFor struct {
	Participant    string         `json:"participant"`
	OriginalAmount *amount.Amount `json:"originalAmount,omitempty"`
	Shares         int            `json:"shares"`
}

// ModifyExpenseFormDocument describes an uploaded document attached to an expense.
type ModifyExpenseFormDocument struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
}

// ModifyExpenseForm contains editable expense fields.
type ModifyExpenseForm struct {
	ExpenseDate                 time.Time                   `json:"expenseDate"`
	Title                       string                      `json:"title"`
	CategoryID                  int                         `json:"category"`
	Amount                      amount.Amount               `json:"amount"`
	OriginalAmount              *amount.Amount              `json:"originalAmount,omitempty"`
	OriginalCurrency            *string                     `json:"originalCurrency,omitempty"`
	ConversionRate              *decimal.Decimal            `json:"conversionRate,omitempty"`
	PaidBy                      string                      `json:"paidBy"`
	PaidFor                     []ModifyExpenseFormPaidFor  `json:"paidFor"`
	SplitMode                   model.SplitMode             `json:"splitMode"`
	SaveDefaultSplittingOptions bool                        `json:"saveDefaultSplittingOptions"`
	IsReimbursement             bool                        `json:"isReimbursement"`
	Documents                   []ModifyExpenseFormDocument `json:"documents,omitempty"`
	Notes                       *string                     `json:"notes,omitempty"`
	RecurrenceRule              model.RecurrenceRule        `json:"recurrenceRule"`
}

// CreateExpenseRequest is the input for the create expense endpoint.
type CreateExpenseRequest struct {
	GroupID       string            `json:"groupId"`
	FormValues    ModifyExpenseForm `json:"expenseFormValues"`
	ParticipantID *string           `json:"participantId,omitempty"`
}

// ApplyDefaults applies default split and recurrence values expected by the API.
func (receiver *CreateExpenseRequest) ApplyDefaults() {
	if receiver.FormValues.SplitMode == "" {
		receiver.FormValues.SplitMode = model.SplitModeEvenly
	}
	if receiver.FormValues.RecurrenceRule == "" {
		receiver.FormValues.RecurrenceRule = model.RecurrenceRuleNone
	}
}

// CreateExpenseResponse contains the identifier of the created expense.
type CreateExpenseResponse struct {
	ExpenseID string `json:"expenseId"`
}
