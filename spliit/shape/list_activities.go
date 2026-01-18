package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

type ListActivitiesRequest struct {
	GroupID string `json:"groupId"`
	Cursor  *uint  `json:"cursor,omitempty"`
	Limit   *uint  `json:"limit,omitempty"`
}

type ListActivitiesResponse struct {
	Activities []model.Activity `json:"activities"`
	HasMore    bool             `json:"hasMore"`
	NextCursor uint             `json:"nextCursor"`
}
