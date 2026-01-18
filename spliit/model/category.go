package model

type Category struct {
	ID       int        `json:"id"`
	Grouping string     `json:"grouping"`
	Name     string     `json:"name"`
	Expenses []*Expense `json:"Expense"`
}
