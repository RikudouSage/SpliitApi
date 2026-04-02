package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// ListBalances describes the groups.balances.list endpoint.
type ListBalances struct {
}

// Name returns the remote endpoint name.
func (receiver *ListBalances) Name() string {
	return "groups.balances.list"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *ListBalances) InputShape() shape.ListBalancesRequest {
	return shape.ListBalancesRequest{}
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *ListBalances) OutputShape() shape.ListBalancesResponse {
	return shape.ListBalancesResponse{}
}

// Mutates reports that the endpoint is read-only.
func (receiver *ListBalances) Mutates() bool {
	return false
}
