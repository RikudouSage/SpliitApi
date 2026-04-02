package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// CreateGroup describes the groups.create endpoint.
type CreateGroup struct {
}

// Name returns the remote endpoint name.
func (receiver *CreateGroup) Name() string {
	return "groups.create"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *CreateGroup) InputShape() shape.CreateGroupRequest {
	return shape.CreateGroupRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *CreateGroup) OutputShape() shape.CreateGroupResponse {
	return shape.CreateGroupResponse{}
}

// Mutates reports that the endpoint changes remote state.
func (receiver *CreateGroup) Mutates() bool {
	return true
}
