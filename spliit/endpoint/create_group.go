package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type CreateGroup struct {
}

func (receiver *CreateGroup) Name() string {
	return "groups.create"
}

func (receiver *CreateGroup) InputShape() shape.CreateGroupRequest {
	return shape.CreateGroupRequest{}
}

func (receiver *CreateGroup) OutputShape() shape.CreateGroupResponse {
	return shape.CreateGroupResponse{}
}

func (receiver *CreateGroup) Mutates() bool {
	return true
}
