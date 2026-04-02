package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

// ListGroupsRequest is the input for the list groups endpoint.
type ListGroupsRequest struct {
	GroupIDs []string `json:"groupIds"`
}

// ListGroupsResponse contains the requested groups.
type ListGroupsResponse struct {
	Groups []model.Group `json:"groups"`
}
