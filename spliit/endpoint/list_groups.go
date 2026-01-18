package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type ListGroups struct {
}

func (receiver *ListGroups) Name() string {
	return "groups.list"
}

func (receiver *ListGroups) InputShape() shape.ListGroupsRequest {
	return shape.ListGroupsRequest{}
}

func (receiver *ListGroups) OutputShape() shape.ListGroupsResponse {
	return shape.ListGroupsResponse{}
}

func (receiver *ListGroups) Mutates() bool {
	return false
}
