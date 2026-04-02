package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// ListActivities describes the groups.activities.list endpoint.
type ListActivities struct {
}

// Name returns the remote endpoint name.
func (receiver *ListActivities) Name() string {
	return "groups.activities.list"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *ListActivities) InputShape() shape.ListActivitiesRequest {
	return shape.ListActivitiesRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *ListActivities) OutputShape() shape.ListActivitiesResponse {
	return shape.ListActivitiesResponse{}
}

// Mutates reports that the endpoint is read-only.
func (receiver *ListActivities) Mutates() bool {
	return false
}
