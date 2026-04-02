package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

// ListCategoriesRequest is the input for the list categories endpoint.
type ListCategoriesRequest = *NilRequest

// ListCategoriesResponse contains all available expense categories.
type ListCategoriesResponse struct {
	Categories []model.Category `json:"categories"`
}
