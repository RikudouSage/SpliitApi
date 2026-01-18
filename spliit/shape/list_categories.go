package shape

import "go.chrastecky.dev/spliit-api/spliit/model"

type ListCategoriesRequest = *NilRequest

type ListCategoriesResponse struct {
	Categories []model.Category `json:"categories"`
}
