package endpoint

import "go.chrastecky.dev/spliit-api/spliit/shape"

type ListCategories struct {
}

func (receiver *ListCategories) Name() string {
	return "categories.list"
}

func (receiver *ListCategories) InputShape() shape.ListCategoriesRequest {
	return nil
}

func (receiver *ListCategories) OutputShape() shape.ListCategoriesResponse {
	return shape.ListCategoriesResponse{}
}

func (receiver *ListCategories) Mutates() bool {
	return false
}
