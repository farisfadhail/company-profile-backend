package request

type ParentCategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}