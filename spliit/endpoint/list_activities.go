package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type ListActivities struct {
}

func (receiver *ListActivities) Name() string {
	return "groups.activities.list"
}

func (receiver *ListActivities) InputShape() shape.ListActivitiesRequest {
	return shape.ListActivitiesRequest{}
}

func (receiver *ListActivities) OutputShape() shape.ListActivitiesResponse {
	return shape.ListActivitiesResponse{}
}

func (receiver *ListActivities) Mutates() bool {
	return false
}
