package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// GetGroupDetails describes the groups.getDetails endpoint.
type GetGroupDetails struct {
}

// Name returns the remote endpoint name.
func (receiver *GetGroupDetails) Name() string {
	return "groups.getDetails"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *GetGroupDetails) InputShape() shape.GetGroupDetailsRequest {
	return shape.GetGroupDetailsRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *GetGroupDetails) OutputShape() shape.GetGroupDetailsResponse {
	return shape.GetGroupDetailsResponse{}
}

// Mutates reports that the endpoint is read-only.
func (receiver *GetGroupDetails) Mutates() bool {
	return false
}
