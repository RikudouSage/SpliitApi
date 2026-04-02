package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// GetGroup describes the groups.get endpoint.
type GetGroup struct {
}

// Name returns the remote endpoint name.
func (receiver *GetGroup) Name() string {
	return "groups.get"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *GetGroup) InputShape() shape.GetGroupRequest {
	return shape.GetGroupRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *GetGroup) OutputShape() shape.GetGroupResponse {
	return shape.GetGroupResponse{}
}

// Mutates reports that the endpoint is read-only.
func (receiver *GetGroup) Mutates() bool {
	return false
}
