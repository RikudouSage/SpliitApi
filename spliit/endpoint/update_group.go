package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type UpdateGroup struct{}

func (receiver *UpdateGroup) Name() string {
	return "groups.update"
}

func (receiver *UpdateGroup) InputShape() shape.UpdateGroupRequest {
	return shape.UpdateGroupRequest{}
}

func (receiver *UpdateGroup) OutputShape() shape.UpdateGroupResponse {
	return nil
}

func (receiver *UpdateGroup) Mutates() bool {
	return true
}
