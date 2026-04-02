package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

// ListActivitiesRequest is the input for the list activities endpoint.
type ListActivitiesRequest struct {
	GroupID string `json:"groupId"`
	Cursor  *uint  `json:"cursor,omitempty"`
	Limit   *uint  `json:"limit,omitempty"`
}

// ListActivitiesResponse contains a page of group activities.
type ListActivitiesResponse struct {
	Activities []model.Activity `json:"activities"`
	HasMore    bool             `json:"hasMore"`
	NextCursor uint             `json:"nextCursor"`
}
