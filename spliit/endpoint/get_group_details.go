package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type GetGroupDetails struct {
}

func (receiver *GetGroupDetails) Name() string {
	return "groups.getDetails"
}

func (receiver *GetGroupDetails) InputShape() shape.GetGroupDetailsRequest {
	return shape.GetGroupDetailsRequest{}
}

func (receiver *GetGroupDetails) OutputShape() shape.GetGroupDetailsResponse {
	return shape.GetGroupDetailsResponse{}
}

func (receiver *GetGroupDetails) Mutates() bool {
	return false
}
