package model

import "time"

type Group struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Information  *string        `json:"information"`
	Currency     string         `json:"currency"`
	CurrencyCode *string        `json:"currencyCode"`
	Participants []*Participant `json:"participants"`
	Expenses     []*Expense     `json:"expenses"`
	Activities   []*Activity    `json:"activities"`
	CreatedAt    time.Time      `json:"createdAt"`

	// list groups endpoint
	Counts map[string]uint `json:"_count"`
}
