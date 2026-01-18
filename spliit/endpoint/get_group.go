package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type GetGroup struct {
}

func (receiver *GetGroup) Name() string {
	return "groups.get"
}

func (receiver *GetGroup) InputShape() shape.GetGroupRequest {
	return shape.GetGroupRequest{}
}

func (receiver *GetGroup) OutputShape() shape.GetGroupResponse {
	return shape.GetGroupResponse{}
}

func (receiver *GetGroup) Mutates() bool {
	return false
}
