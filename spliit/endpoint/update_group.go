package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// UpdateGroup describes the groups.update endpoint.
type UpdateGroup struct{}

// Name returns the remote endpoint name.
func (receiver *UpdateGroup) Name() string {
	return "groups.update"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *UpdateGroup) InputShape() shape.UpdateGroupRequest {
	return shape.UpdateGroupRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *UpdateGroup) OutputShape() shape.UpdateGroupResponse {
	return nil
}

// Mutates reports that the endpoint changes remote state.
func (receiver *UpdateGroup) Mutates() bool {
	return true
}
