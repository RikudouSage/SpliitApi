package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// ListGroups describes the groups.list endpoint.
type ListGroups struct {
}

// Name returns the remote endpoint name.
func (receiver *ListGroups) Name() string {
	return "groups.list"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *ListGroups) InputShape() shape.ListGroupsRequest {
	return shape.ListGroupsRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *ListGroups) OutputShape() shape.ListGroupsResponse {
	return shape.ListGroupsResponse{}
}

// Mutates reports that the endpoint is read-only.
func (receiver *ListGroups) Mutates() bool {
	return false
}
