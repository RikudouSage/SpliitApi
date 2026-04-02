package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

// GetGroupRequest is the input for the get group endpoint.
type GetGroupRequest struct {
	GroupID string `json:"groupId"`
}

// GetGroupResponse contains the requested group.
type GetGroupResponse struct {
	Group *model.Group `json:"group"`
}
