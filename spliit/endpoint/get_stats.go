package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// GetStats describes the groups.stats.get endpoint.
type GetStats struct{}

// Name returns the remote endpoint name.
func (receiver *GetStats) Name() string {
	return "groups.stats.get"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *GetStats) InputShape() shape.GetStatsRequest {
	return shape.GetStatsRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *GetStats) OutputShape() shape.GetStatsResponse {
	return shape.GetStatsResponse{}
}

// Mutates reports that the endpoint is read-only.
func (receiver *GetStats) Mutates() bool {
	return false
}
