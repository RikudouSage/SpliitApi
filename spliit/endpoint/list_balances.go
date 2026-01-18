package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type ListBalances struct {
}

func (receiver *ListBalances) Name() string {
	return "groups.balances.list"
}

func (receiver *ListBalances) InputShape() shape.ListBalancesRequest {
	return shape.ListBalancesRequest{}
}

func (receiver *ListBalances) OutputShape() shape.ListBalancesResponse {
	return shape.ListBalancesResponse{}
}

func (receiver *ListBalances) Mutates() bool {
	return false
}
