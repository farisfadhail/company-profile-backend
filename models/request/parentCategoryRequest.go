package request

type ParentCategoryRequest struct {
	Name string `json:"name" validate:"required"`
	Slug string `json:"slug"`
}