package shape

// ModifyGroupParticipant describes a participant submitted in group create and update forms.
type ModifyGroupParticipant struct {
	ID   *string `json:"id,omitempty"`
	Name string  `json:"name"`
}

// ModifyGroupForm contains editable group fields.
type ModifyGroupForm struct {
	Name         string                   `json:"name"`
	Information  *string                  `json:"information,omitempty"`
	Currency     string                   `json:"currency"`
	CurrencyCode *string                  `json:"currencyCode,omitempty"`
	Participants []ModifyGroupParticipant `json:"participants"`
}

// CreateGroupRequest is the input for the create group endpoint.
type CreateGroupRequest struct {
	FormValues ModifyGroupForm `json:"groupFormValues"`
}

// CreateGroupResponse contains the identifier of the created group.
type CreateGroupResponse struct {
	GroupID string `json:"groupId"`
}
