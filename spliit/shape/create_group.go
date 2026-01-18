package shape

type ModifyGroupParticipant struct {
	ID   *string `json:"id,omitempty"`
	Name string  `json:"name"`
}

type ModifyGroupForm struct {
	Name         string                   `json:"name"`
	Information  *string                  `json:"information,omitempty"`
	Currency     string                   `json:"currency"`
	CurrencyCode *string                  `json:"currency_code,omitempty"`
	Participants []ModifyGroupParticipant `json:"participants"`
}

type CreateGroupRequest struct {
	FormValues ModifyGroupForm `json:"groupFormValues"`
}

type CreateGroupResponse struct {
	GroupID string `json:"groupId"`
}
