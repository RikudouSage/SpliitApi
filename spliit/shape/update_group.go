package shape

type UpdateGroupRequest struct {
	GroupID       string          `json:"groupId"`
	FormValues    ModifyGroupForm `json:"groupFormValues"`
	ParticipantID *string         `json:"participantId,omitempty"`
}

type UpdateGroupResponse = *string
