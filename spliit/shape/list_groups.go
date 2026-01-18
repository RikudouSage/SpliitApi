package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

type ListGroupsRequest struct {
	GroupIDs []string `json:"groupIds"`
}

type ListGroupsResponse struct {
	Groups []model.Group `json:"groups"`
}
