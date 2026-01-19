package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

type GetGroupRequest struct {
	GroupID string `json:"groupId"`
}

type GetGroupResponse struct {
	Group *model.Group `json:"group"`
}
