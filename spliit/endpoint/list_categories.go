package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

// ListCategories describes the categories.list endpoint.
type ListCategories struct {
}

// Name returns the remote endpoint name.
func (receiver *ListCategories) Name() string {
	return "categories.list"
}

// InputShape returns the request DTO shape for the endpoint.
func (receiver *ListCategories) InputShape() shape.ListCategoriesRequest {
	return nil
}

// OutputShape returns the response DTO shape for the endpoint.
func (receiver *ListCategories) OutputShape() shape.ListCategoriesResponse {
	return shape.ListCategoriesResponse{}
}

// Mutates reports that the endpoint is read-only.
func (receiver *ListCategories) Mutates() bool {
	return false
}
