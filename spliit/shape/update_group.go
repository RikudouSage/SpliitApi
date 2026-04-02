package shape

// UpdateGroupRequest is the input for the update group endpoint.
type UpdateGroupRequest struct {
	GroupID       string          `json:"groupId"`
	FormValues    ModifyGroupForm `json:"groupFormValues"`
	ParticipantID *string         `json:"participantId,omitempty"`
}

// UpdateGroupResponse is the response returned by the update group endpoint.
type UpdateGroupResponse = *string
