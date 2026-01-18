package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type GetStats struct{}

func (receiver *GetStats) Name() string {
	return "groups.stats.get"
}

func (receiver *GetStats) InputShape() shape.GetStatsRequest {
	return shape.GetStatsRequest{}
}

func (receiver *GetStats) OutputShape() shape.GetStatsResponse {
	return shape.GetStatsResponse{}
}

func (receiver *GetStats) Mutates() bool {
	return false
}
